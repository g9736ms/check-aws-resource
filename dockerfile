# Golang 이미지를 베이스로 사용합니다.
FROM golang:latest AS builder

# 작업 디렉토리를 생성합니다.
WORKDIR /app

# 소스 코드를 복사합니다.
COPY . .

RUN go mod init github.com/g9736ms/check-aws-resource

RUN go mod init github.com/aws/aws-sdk-go/aws

# 의존성 패키지를 설치합니다.
RUN go mod download

# 소스 코드를 빌드합니다.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o check-aws-resource ./cmd/main.go

# 빌드 결과물을 실행할 이미지를 생성합니다.
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .

# 컨테이너가 실행될 때 실행할 명령어를 지정합니다.
CMD ["./check-aws-resource"]