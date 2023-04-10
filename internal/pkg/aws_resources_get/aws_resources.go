package aws_resources

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// 다른 리소스 들도 검색 및 추가 할 수 있도록 바꿈
type AWSUnusedResources interface {
	GetUnusedEIPs() ([]string, error)
	GetUnusedEBSs() ([]string, error)
}

type AWSClient struct {
	svcEc2 *ec2.EC2
}

// 중복 되는 인증 부분 하나로 해버리자 !!
func NewAWSClient(region string) (*AWSClient, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return nil, err
	}

	return &AWSClient{
		svcEc2: ec2.New(sess),
	}, nil
}
