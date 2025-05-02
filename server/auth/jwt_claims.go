package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sol-armada/commander/configs"
	"github.com/sol-armada/commander/member"
)

type JwtClaims struct {
	Member *member.Member `json:"member,omitempty"`
	jwt.RegisteredClaims
}

func GenerateJWT(member *member.Member) (string, error) {
	expirationTime := time.Now().Add(5 * time.Hour)

	claims := &JwtClaims{
		Member: member,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    configs.DiscordClientId,
			Audience:  []string{configs.DiscordClientId},
			Subject:   member.Id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(configs.JwtKey))
}
