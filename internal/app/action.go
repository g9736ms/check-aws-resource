package action

import (
	"../../internal/pkg/aws/get"
	"../../internal/pkg/env_read"
	"../../internal/pkg/slack_send"
	"fmt"
	"log"
	"strings"
)

// TODO  여기에다가 두개 합쳐서 매인에서 뿅 쏠거
func MainAction() error {
	unusedResources, err := AWSResult()
	if err != nil {
		log.Fatalf("Error in AWSResult: %v", err)
	}

	err = SendResult(unusedResources)

	return nil
}

// TODO 슬렉 쏘는건 만들어 놨으나 텐스트 값을 받아와서 같이 쏴야 함
func SendResult(unusedResources map[string][]string) error {
	// 값을 가져와 사용
	webhookurl := env_reader.GetEnv("webhookURL")
	fmt.Println(webhookurl)
	sender := slack_sender.HTTPSlackSender{} // 인터페이스 선
	var messageBuilder strings.Builder
	messageBuilder.WriteString("사용 중이지 않은 AWS리 소스를 알려드립니다.\n")

	for resourceType, resourceList := range unusedResources {
		messageBuilder.WriteString(resourceType + ":\n")
		for _, resource := range resourceList {
			messageBuilder.WriteString(fmt.Sprintf("- %s\n", resource))
		}
	}

	infoMessage := messageBuilder.String()

	err := sender.SendMessage(webhookurl, infoMessage)
	if err != nil {
		log.Fatal(err)
	}

	return err

}

// TODO 여기는 AWS 결과를 받아오고 리턴을 해줘야 함 리턴 부분 아직 안만들었음
func AWSResult() (map[string][]string, error) {
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

	unusedEIPs, err := awsClient.GetUnusedEIPs()
	if err != nil {
		log.Fatalf("Error getting unused ELBs: %v", err)
	}

	result := map[string][]string{
		"UnusedEBSs": unusedEBSs,
		"UnusedEIPs": unusedEIPs,
	}

	return result, err
}
