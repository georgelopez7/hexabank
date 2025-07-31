package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DiscordClient struct {
	webhookURL string
}

func NewDiscordClient(webhookURL string) *DiscordClient {
	return &DiscordClient{
		webhookURL: webhookURL,
	}
}

type DiscordMessage struct {
	Content string `json:"content"`
}

func (d *DiscordClient) SendMessage(message string) error {
	payload := DiscordMessage{Content: message}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	resp, err := http.Post(d.webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	return nil
}
