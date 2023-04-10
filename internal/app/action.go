package action

import (
	"../../internal/pkg/aws_resources_get"
	"../../internal/pkg/env_read"
	"../../internal/pkg/slack_send"
	"../../lib/slack/template"
	"fmt"
	"log"
)

// TODO  여기에다가 두개 합쳐서 매인에서 뿅 쏠거
func MainAction() error {

	return nil
}

// TODO 슬렉 쏘는건 만들어 놨으나 텐스트 값을 받아와서 같이 쏴야 함
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

// TODO 여기는 AWS 결과를 받아오고 리턴을 해줘야 함 리턴 부분 아직 안만들었음
func AWSResult() error {
	// 세션 중복 되서 통합적으로 해줄려고 만듬
	awsClient, err := aws_resources.NewAWSClient("ap-northeast-2")
	if err != nil {
		log.Fatalf("Error creating AWS client: %v", err)
	}
	//TODO 인터페이스 에 있는 저거 호출해서 쭈루룩 받아오게 하자
	unusedEBSs, err := awsClient.GetUnusedEBSs()
	if err != nil {
		log.Fatalf("Error getting unused ELBs: %v", err)
	}
	// 보이게 만들 부분 주석 처리 요망
	fmt.Printf("Unused EBSs: %v\n", unusedEBSs)

	unusedEIPs, err := awsClient.GetUnusedEIPs()
	if err != nil {
		log.Fatalf("Error getting unused ELBs: %v", err)
	}
	// 보이게 만들 부분 주석 처리 요망
	fmt.Printf("Unused EBSs: %v\n", unusedEIPs)

	return err
}
