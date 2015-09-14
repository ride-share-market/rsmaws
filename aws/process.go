package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func BuildVpc(vpcRegion string) (string, error) {

	svc := ec2.New(&aws.Config{
		Region: aws.String(vpcRegion),
	})

	awscvpc := new(AwsCreateVpc)
	vpcId, err := CreateVpc(awscvpc, svc)
	if err != nil {
		return "", err
	}

	return vpcId, nil

}
