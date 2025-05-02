package auth

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/sol-armada/commander/configs"
	"github.com/sol-armada/commander/member"
	"github.com/sol-armada/sol-bot/stores"
)

func Login(s *stores.Client, code string) (string, error) {
	authLogger := slog.With("endpoint", "authenticate")

	data := url.Values{}

	data.Set("client_id", configs.DiscordClientId)
	data.Set("client_secret", configs.DiscordClientSecret)
	data.Set("redirect_uri", configs.DiscordRedirectUri)
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)

	// authenticate with discord
	req, err := http.NewRequest("POST", "https://discord.com/api/v10/oauth2/token", strings.NewReader(data.Encode()))
	if err != nil {
		authLogger.Error("failed to create request", "err", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		authLogger.Error("failed to do request", "err", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return "", errors.New("invalid_grant")
	}

	if resp.StatusCode == http.StatusBadRequest {
		errorMessage, _ := io.ReadAll(resp.Body)
		type ErrorMessage struct {
			ErrorType   string `json:"error"`
			Description string `json:"error_description"`
		}
		errMsg := ErrorMessage{}
		if err := json.Unmarshal(errorMessage, &errMsg); err != nil {
			return "", err
		}
		if errMsg.ErrorType == "invalid_grant" {
			return "", errors.New(errMsg.Description)
		}

		return "", errors.New(errMsg.Description)
	}

	access := map[string]any{}
	if err := json.NewDecoder(resp.Body).Decode(&access); err != nil {
		authLogger.Error("failed to decode response", "err", err)
		return "", err
	}

	if errStr, ok := access["error"].(string); ok {
		authLogger.Error("failed to authenticate", "err", err)
		return "", errors.New(errStr)
	}

	accessToken := access["access_token"].(string)

	authLogger.Debug("access token", "access_token", accessToken)

	// get the access token
	req, err = http.NewRequest("GET", "https://discord.com/api/users/@me", nil)
	if err != nil {
		authLogger.Error("failed to create request", "err", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err = client.Do(req)
	if err != nil {
		authLogger.Error("failed to do request", "err", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			return "", err
		}

		authLogger.Error("failed to authenticate", "err", err)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		authLogger.Error("failed to read body", "err", err)
		return "", err
	}

	userMap := map[string]any{}
	if err := json.Unmarshal(body, &userMap); err != nil {
		authLogger.Error("failed to unmarshal body", "err", err)
		return "", err
	}

	if userMap["id"] == nil {
		authLogger.Error("failed to get user id", "err", err)
		return "", errors.New("failed to get user id")
	}

	// solMember, err := solmembers.Get(userMap["id"].(string))
	// if err != nil {
	// 	authLogger.Error("failed to get sol member", "err", err)
	// 	return "", err
	// }
	// if solMember == nil {
	// 	authLogger.Error("failed to get sol member", "err", err)
	// 	return "", errors.New("failed to get sol member")
	// }
	m, err := member.Get(s, userMap["id"].(string))
	if err != nil {
		authLogger.Error("failed to get member", "err", err)
		return "", err
	}
	if m == nil {
		authLogger.Error("failed to get member", "err", err)
		return "", errors.New("failed to get member")
	}

	return GenerateJWT(m)
}
