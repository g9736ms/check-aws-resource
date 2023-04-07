package main

import (
	"../internal/env_read"
	"fmt"
)

func main() {
	// Load environment variables from file
	err := env_reader.SetEnv("/tmp/test.txt")
	if err != nil {
		fmt.Println(err)
	}

	// Get the value of an environment variable
	dbHost := env_reader.GetEnv("DB_HOST")
	fmt.Println("DB_HOST:", dbHost)

}

// 파일을 읽어서 키를 가지고 올 수 있는 것을 만들자

// aws sdk 로 쿼리 할 수 있는 부분을 만들자

// 슬렉으로 보낼 부분에 대해서 토큰 값을 읽어올 수 있는 부분을 만들자

// 받아온 값들을 슬렉으로 보낼 부분을 만들자

// 안쓰는 리소스들을 찾고싶습니다 .. 금요일날 오전에 8시쯤
