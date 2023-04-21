package aws_resources

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (a AWSClient) GetUnusedEIPs() ([]string, error) {
	input := &ec2.DescribeAddressesInput{}
	result, err := a.svcEc2.DescribeAddresses(input)

	if err != nil {
		return nil, err
	}

	var unusedEIPs []string
	for _, address := range result.Addresses {
		if address.PrivateIpAddress == nil {
			unusedEIPs = append(unusedEIPs, *address.PublicIp)
		}
	}

	return unusedEIPs, nil
}
