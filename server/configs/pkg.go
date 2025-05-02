package configs

import (
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/rs/xid"
)

var (
	Origins             = []string{}
	DiscordClientId     string
	DiscordClientSecret string
	DiscordRedirectUri  string

	CacheHost string
	CachePort int

	Debug       bool
	LogPath     string
	LogLevel    string
	LogFile     string
	EchoLogFile string

	JwtKey string

	MongoHost           string
	MongoPort           int
	MongoUser           string
	MongoPass           string
	MongoDB             string
	MongoReplicaSetName string
)

func Load() {
	originsRaw := flag.String("origins", "*", "Allowed origins for CORS")
	flag.StringVar(&DiscordClientId, "discord-client-id", "", "Discord client ID")
	flag.StringVar(&DiscordClientSecret, "discord-client-secret", "", "Discord client secret")
	flag.StringVar(&DiscordRedirectUri, "discord-redirect-uri", "", "Discord redirect URI")

	flag.BoolVar(&Debug, "debug", false, "Enable debug mode")
	flag.StringVar(&LogPath, "log-path", "./", "Path to log file")
	flag.StringVar(&LogLevel, "log-level", "info", "Log level (debug, info, warn, error)")
	flag.StringVar(&LogFile, "log-file", "logs", "Log file name")
	flag.StringVar(&EchoLogFile, "echo-log-file", "echo-logs", "Path to log file")

	flag.StringVar(&JwtKey, "jwt-key", xid.New().String(), "JWT key for authentication")

	flag.StringVar(&MongoHost, "mongo-host", "localhost", "MongoDB host")
	flag.IntVar(&MongoPort, "mongo-port", 27017, "MongoDB port")
	flag.StringVar(&MongoUser, "mongo-user", "", "MongoDB username")
	flag.StringVar(&MongoPass, "mongo-pass", "", "MongoDB password")
	flag.StringVar(&MongoDB, "mongo-db", "", "MongoDB database name")
	flag.StringVar(&MongoReplicaSetName, "mongo-replica-set-name", "", "MongoDB replica set name")

	flag.Parse()

	if *originsRaw != "" {
		Origins = strings.Split(*originsRaw, ",")
	}

	// overwrite with env vars
	if envOrigins := GetEnv("ORIGINS"); envOrigins != "" {
		Origins = strings.Split(envOrigins, ",")
	}
	if envDiscordClientId := GetEnv("DISCORD_CLIENT_ID"); envDiscordClientId != "" {
		DiscordClientId = envDiscordClientId
	}
	if envDiscordClientSecret := GetEnv("DISCORD_CLIENT_SECRET"); envDiscordClientSecret != "" {
		DiscordClientSecret = envDiscordClientSecret
	}
	if envDiscordRedirectUri := GetEnv("DISCORD_REDIRECT_URI"); envDiscordRedirectUri != "" {
		DiscordRedirectUri = envDiscordRedirectUri
	}
	if envCacheHost := GetEnv("CACHE_HOST"); envCacheHost != "" {
		CacheHost = envCacheHost
	}
	if envCachePort := GetEnv("CACHE_PORT"); envCachePort != "" {
		if port, err := strconv.Atoi(envCachePort); err == nil {
			CachePort = port
		}
	}
	if envDebug := GetEnv("DEBUG"); envDebug != "" {
		if envDebug == "true" {
			Debug = true
		} else {
			Debug = false
		}
	}
	if envLogPath := GetEnv("LOG_PATH"); envLogPath != "" {
		LogPath = envLogPath
	}
	if envLogLevel := GetEnv("LOG_LEVEL"); envLogLevel != "" {
		LogLevel = envLogLevel
	}
	if envLogFile := GetEnv("LOG_FILE"); envLogFile != "" {
		LogFile = envLogFile
	}
	if envEchoLogFile := GetEnv("ECHO_LOG_FILE"); envEchoLogFile != "" {
		EchoLogFile = envEchoLogFile
	}
	if envJwtKey := GetEnv("JWT_KEY"); envJwtKey != "" {
		JwtKey = envJwtKey
	}
	if envMongoHost := GetEnv("MONGO_HOST"); envMongoHost != "" {
		MongoHost = envMongoHost
	}
	if envMongoPort := GetEnv("MONGO_PORT"); envMongoPort != "" {
		if port, err := strconv.Atoi(envMongoPort); err == nil {
			MongoPort = port
		}
	}
	if envMongoUser := GetEnv("MONGO_USER"); envMongoUser != "" {
		MongoUser = envMongoUser
	}
	if envMongoPass := GetEnv("MONGO_PASS"); envMongoPass != "" {
		MongoPass = envMongoPass
	}
	if envMongoDB := GetEnv("MONGO_DB"); envMongoDB != "" {
		MongoDB = envMongoDB
	}
	if envMongoReplicaSetName := GetEnv("MONGO_REPLICA_SET_NAME"); envMongoReplicaSetName != "" {
		MongoReplicaSetName = envMongoReplicaSetName
	}
}

func GetEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return ""
}
