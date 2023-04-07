package main

import (
	"../internal/env_read"
	"../internal/slack_sender"
	"fmt"
	"log"
)

func main() {
	// 변수 파일 읽기
	err := env_reader.SetEnv("/tmp/test.txt")
	if err != nil {
		fmt.Println(err)
	}

	// 값을 가져와 사용
	webhookurl := env_reader.GetEnv("webhookURL")
	fmt.Println(webhookurl)
	sender := slack_sender.HTTPSlackSender{} // 인터페이스 선
	message := "Hello, Slack!"

	err = sender.SendMessage(webhookurl, message)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Message sent to Slack!")

}

// 파일을 읽어서 키를 가지고 올 수 있는 것을 만들자

// aws sdk 로 쿼리 할 수 있는 부분을 만들자

// 슬렉으로 보낼 부분에 대해서 토큰 값을 읽어올 수 있는 부분을 만들자

// 받아온 값들을 슬렉으로 보낼 부분을 만들자

// 안쓰는 리소스들을 찾고싶습니다 .. 금요일날 오전에 8시쯤
