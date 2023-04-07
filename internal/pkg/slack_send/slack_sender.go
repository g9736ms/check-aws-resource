package slack_sender

type SlackSender interface {
	SendMessage(webhookURL string, message string) error
}
