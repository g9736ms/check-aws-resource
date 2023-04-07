package aws_resources

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (a AWSEc2) GetUnusedEBSs() ([]string, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"), // 필요한 리전으로 변경하세요.
	}))

	svc := ec2.New(sess)
	input := &ec2.DescribeVolumesInput{}
	result, err := svc.DescribeVolumes(input)

	if err != nil {
		return nil, err
	}

	var unusedEBSs []string
	for _, volume := range result.Volumes {
		if *volume.State == "available" {
			unusedEBSs = append(unusedEBSs, *volume.VolumeId)
		}
	}

	return unusedEBSs, nil
}
