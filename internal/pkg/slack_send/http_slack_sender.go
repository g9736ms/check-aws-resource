// http_slack_sender.go
package slack_sender

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type HTTPSlackSender struct{}

type slackPayload struct {
	Text string `json:"text"`
}

func (s HTTPSlackSender) SendMessage(webhookURL string, message string) error {
	payload := slackPayload{
		Text: message,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to send message to Slack")
	}

	return nil
}
