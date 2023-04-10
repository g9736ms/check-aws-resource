package aws_resources

// 2버전으로 해야합니다 .
import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (a AWSClient) GetUnusedEBSs() ([]string, error) {
	input := &ec2.DescribeVolumesInput{}
	result, err := a.svcEc2.DescribeVolumes(input)

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
