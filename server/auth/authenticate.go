package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/golang-jwt/jwt/v5"
	middleware "github.com/oapi-codegen/echo-middleware"
	"github.com/sol-armada/commander/configs"
)

func NewAuthenticator() openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return Authenticate(ctx, input)
	}
}

func Authenticate(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	if input.SecuritySchemeName != "bearerAuth" {
		return fmt.Errorf("unexpected security scheme: %s", input.SecuritySchemeName)
	}

	tokenString, err := getTokenFromHeader(input.RequestValidationInput.Request)
	if err != nil {
		return fmt.Errorf("failed to get token from header: %v", err)
	}

	token, err := GetToken(tokenString)
	if err != nil {
		return fmt.Errorf("failed to get token: %v", err)
	}

	eCtx := middleware.GetEchoContext(ctx)
	eCtx.Set("member", *token.Claims.(*JwtClaims).Member)

	return nil
}

func GetToken(t string) (*jwt.Token, error) {
	parsedToken, err := jwt.ParseWithClaims(t, &JwtClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.JwtKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if err := checkTokenClaims(parsedToken); err != nil {
		return nil, fmt.Errorf("failed to check token claims: %v", err)
	}

	return parsedToken, nil
}

func checkTokenClaims(t *jwt.Token) error {
	_, ok := t.Claims.(*JwtClaims)
	if !ok {
		return fmt.Errorf("failed to parse claims")
	}

	return nil
}

func getTokenFromHeader(req *http.Request) (string, error) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("missing Authorization header")
	}

	prefix := "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", fmt.Errorf("invalid Authorization header format")
	}

	return strings.TrimPrefix(authHeader, prefix), nil
}
