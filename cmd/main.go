package main

import (
	"../internal/pkg/aws_resources_get"
	"../internal/pkg/env_read"
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
	var err error
	// env 세팅
	if err = env_reader.SetEnv(args); err != nil {
		fmt.Println(err)
	}

	/*
		// action 을 따로 빼버리고 send result 여기서 slack 이나 다른 걸로 도 보낼 수 있게 만들 수 있을듯
		if err = action.SendResult(); err != nil {
			fmt.Println(err)
		}
	*/

	if err = test_action(); err != nil {
		fmt.Println(err)
	}

	return err
}

func test_action() error {
	awsClient := aws_resources.AWSEc2{}

	unusedEIPs, err := awsClient.GetUnusedEIPs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Unused EIPs:")
	for _, eip := range unusedEIPs {
		fmt.Println(eip)
	}

	unusedEBSs, err := awsClient.GetUnusedEBSs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nUnused EBSs:")
	for _, ebs := range unusedEBSs {
		fmt.Println(ebs)
	}

	return nil
}

// 파일을 읽어서 키를 가지고 올 수 있는 것을 만들자

// aws sdk 로 쿼리 할 수 있는 부분을 만들자

// 슬렉으로 보낼 부분에 대해서 토큰 값을 읽어올 수 있는 부분을 만들자

// 받아온 값들을 슬렉으로 보낼 부분을 만들자

// 안쓰는 리소스들을 찾고싶습니다 .. 금요일날 오전에 8시쯤
