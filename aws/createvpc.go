package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type CreateVpcIface interface {
	Create(vpcRegion string) (string, error)
}

func CreateVpc(vpc CreateVpcIface, vpcRegion string) (string, error) {
	return vpc.Create(vpcRegion)
}

type AwsCreateVpc struct{}

func (this AwsCreateVpc) Create(vpcRegion string) (string, error) {
	
	svc := ec2.New(&aws.Config{
		Region: aws.String(vpcRegion),
	})

	params := &ec2.CreateVpcInput{
		CidrBlock: aws.String("10.0.0.0/16"),
	}
	
	resp, err := svc.CreateVpc(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return "", err
	}
	
	return *resp.Vpc.VpcId, nil
	
	// Pretty-print the response data.
	//	fmt.Println(resp)
}
