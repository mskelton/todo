package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mskelton/todo/internal/config"
)

type SyncResponse struct {
	FullSync      bool              `json:"full_sync"`
	Projects      []Project         `json:"projects"`
	SyncToken     string            `json:"sync_token"`
	TempIDMapping map[string]uint32 `json:"temp_id_mapping"`
}

func getToken() (string, error) {
	config, err := config.Get()
	if err != nil {
		return "", fmt.Errorf("failed to get config %w", err)
	}

	return config.ApiToken, nil
}

func Sync(syncToken string) (*SyncResponse, error) {
	token, err := getToken()
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"resource_types": []string{"projects"},
		"sync_token":     syncToken,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.todoist.com/sync/v9/sync", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to get create request %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request %w", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	var result SyncResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response %w", err)
	}

	return &result, nil
}
