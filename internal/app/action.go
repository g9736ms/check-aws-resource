package action

import (
	"../../internal/pkg/env_read"
	"../../internal/pkg/slack_send"
	"../../lib/slack/template"
	"fmt"
	"log"
)

func SendResult() error {
	// 값을 가져와 사용
	webhookurl := env_reader.GetEnv("webhookURL")
	fmt.Println(webhookurl)
	sender := slack_sender.HTTPSlackSender{} // 인터페이스 선
	infoMessage := template.InfoMessage("This is an info message.")
	warnMessage := template.WarnMessage("This is a warning message.")

	err := sender.SendMessage(webhookurl, infoMessage)
	if err != nil {
		log.Fatal(err)
	}

	err = sender.SendMessage(webhookurl, warnMessage)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
