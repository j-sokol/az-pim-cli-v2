package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetAccessToken() (string, error) {
	cmd := exec.Command("az", "account", "get-access-token", "--scope=https://management.core.windows.net/.default", "--query", "accessToken", "--output", "tsv")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("failed to get access token: %v", err)
	}

	accessToken := strings.TrimSpace(string(output))
	return accessToken, nil
}

func GetCurrentUserObjectId() (string, error) {
	cmd := exec.Command("az", "ad", "signed-in-user", "show", "--query", "id", "--output", "tsv")
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("failed to get current user object id: %v", err)
	}

	userId := strings.TrimSpace(string(output))
	return userId, nil
}
