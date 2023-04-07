package aws_resources

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AWSEc2 struct{}

func (a AWSEc2) GetUnusedEIPs() ([]string, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"), // 필요한 리전으로 변경하세요.
	}))

	svc := ec2.New(sess)
	input := &ec2.DescribeAddressesInput{}
	result, err := svc.DescribeAddresses(input)

	if err != nil {
		return nil, err
	}

	var unusedEIPs []string
	for _, address := range result.Addresses {
		if address.InstanceId == nil {
			unusedEIPs = append(unusedEIPs, *address.PublicIp)
		}
	}

	return unusedEIPs, nil
}
