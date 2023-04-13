package main

import (
	"../internal/app"
	"../internal/pkg/env_read"
	"fmt"
)

var env_file = "/tmp/test.txt"

func main() {
	// 변수 파일 읽기
	if err := root(env_file); err != nil {
		fmt.Println(err)
	}

	//log.Println("Message sent to Slack!")
}

func root(args string) error {
	var err error
	// env 세팅
	if err = env_reader.SetEnv(args); err != nil {
		fmt.Println(err)
	}

	if err = action.MainAction(); err != nil {
		fmt.Println(err)
	}

	return err
}

func test_action() error {

	return nil
}

// aws sdk 로 쿼리 할 수 있는 부분을 만들자

// 금요일날 오전에 8시쯤 리포트 형태로 실행 되었으면 좋겠음
// 이건 argo workflow으로 하면 좋을 듯 합니다 ㅎㅎ !!

// CI/CD도 만들자 !
// CI 는 워커플로우 사용 예정 그리고 수동으로 워커플로우를 이용할 수 있나 ? 확인해보기
// 이 깃헙을 받아서 빌드 시켜서 우리 레포에 push시 키기
// CD를 위한 부분 변경 하기
