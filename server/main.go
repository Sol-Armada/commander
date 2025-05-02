package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
	slogmulti "github.com/samber/slog-multi"
	"github.com/sol-armada/commander/api"
	"github.com/sol-armada/commander/auth"
	"github.com/sol-armada/commander/configs"
	"github.com/sol-armada/commander/database"
	solmembers "github.com/sol-armada/sol-bot/members"
	"github.com/sol-armada/sol-bot/stores"
	valkeyapi "github.com/valkey-io/valkey-glide/go/api"
)

type Server struct {
	DB *stores.Client
}

var logger *slog.Logger
var ctx context.Context = context.Background()

func init() {
	configs.Load()

	if err := os.MkdirAll(configs.LogPath, os.ModePerm); err != nil {
		slog.Error("failed to create log path", "err", err)
		os.Exit(1)
	}

	logFile, err := os.OpenFile(path.Join(configs.LogPath, configs.LogFile), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("failed to open log file", "err", err)
		os.Exit(1)
	}

	handlerOpts := &slog.HandlerOptions{
		AddSource: true,
	}

	if configs.Debug {
		handlerOpts.Level = slog.LevelDebug
	}

	var logFileHandler slog.Handler = slog.NewTextHandler(logFile, handlerOpts)
	var logConsoleHandler slog.Handler = slog.NewTextHandler(os.Stdout, handlerOpts)

	var fanoutHandler slog.Handler = slogmulti.Fanout(
		logFileHandler,
		logConsoleHandler,
	)

	logger = slog.New(fanoutHandler)
}

func main() {
	opts := &database.Options{
		Host:           configs.MongoHost,
		Port:           configs.MongoPort,
		Username:       configs.MongoUser,
		Password:       configs.MongoPass,
		Database:       configs.MongoDB,
		ReplicaSetName: configs.MongoReplicaSetName,
	}
	dbClient, err := database.New(ctx, opts)
	if err != nil {
		logger.Error("failed to create database client", "err", err)
		os.Exit(1)
	}

	if err := solmembers.Setup(); err != nil {
		logger.Error("failed to setup solmembers", "err", err)
		os.Exit(1)
	}

	cacheConfig := valkeyapi.NewGlideClientConfiguration().WithAddress(
		&valkeyapi.NodeAddress{Host: configs.CacheHost, Port: configs.CachePort},
	)

	c, err := valkeyapi.NewGlideClient(cacheConfig)
	if err != nil {
		logger.Error("failed to create cache client", "err", err)
		os.Exit(1)
	}
	cache := &cache{
		GlideClientCommands: c,
	}
	defer cache.Close()

	echoLogFile, err := os.OpenFile(path.Join(configs.LogPath, configs.EchoLogFile), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Error("failed to open echo log file", "err", err)
		os.Exit(1)
	}
	defer echoLogFile.Close()

	echoLogger := slog.New(slogmulti.Fanout(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
		slog.NewTextHandler(echoLogFile, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	))

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogLatency: true,
		LogError:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			echoLogger.LogAttrs(ctx, slog.LevelInfo, "REQUEST",
				slog.String("uri", v.URI),
				slog.String("method", v.Method),
				slog.Int("status", v.Status),
			)

			if v.Error == nil {
				echoLogger.LogAttrs(ctx, slog.LevelInfo, "REQUEST",
					slog.Int("status", v.Status),
				)
				return nil
			}

			echoLogger.LogAttrs(ctx, slog.LevelError, "REQUEST",
				slog.String("error", v.Error.Error()),
			)
			return nil
		},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     configs.Origins,
		AllowMethods:     []string{http.MethodGet},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	e.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			if req.URL.Path == "/login" || req.URL.Path == "/login/" || req.URL.Path == "/login/*" {
				req.URL.Path = "/"
			}
			return next(c)
		}
	})

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5:      true,
		Root:       "dist",
		Filesystem: http.FS(staticFiles),
	}))

	spec, err := api.GetSwagger()
	if err != nil {
		slog.Error("failed to get swagger spec", "err", err)
		os.Exit(1)
	}

	validator := oapimiddleware.OapiRequestValidatorWithOptions(spec,
		&oapimiddleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: auth.NewAuthenticator(),
			},
		})

	e.Use([]echo.MiddlewareFunc{validator}...)

	server := api.NewServer(logger, dbClient)
	api.RegisterHandlers(e, server)

	go func() {
		e.Logger.Fatal(e.Start(":8080"))
	}()

	// Watch for ctl-c
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	slog.Info("shutting down...")
}
