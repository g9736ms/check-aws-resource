package main

import (
	"../internal/pkg/env_read"
	"../internal/pkg/slack_send"
	"../lib/slack/template"
	"fmt"
	"log"
)

var env_file = "/tmp/test.txt"

func main() {
	// 변수 파일 읽기
	if err := root(env_file); err != nil {
		fmt.Println(err)
	}

	log.Println("Message sent to Slack!")
}

func root(args string) error {
	// env 를 세팅한다.
	if err := env_reader.SetEnv(args); err != nil {
		fmt.Println(err)
	}
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

// 파일을 읽어서 키를 가지고 올 수 있는 것을 만들자

// aws sdk 로 쿼리 할 수 있는 부분을 만들자

// 슬렉으로 보낼 부분에 대해서 토큰 값을 읽어올 수 있는 부분을 만들자

// 받아온 값들을 슬렉으로 보낼 부분을 만들자

// 안쓰는 리소스들을 찾고싶습니다 .. 금요일날 오전에 8시쯤
