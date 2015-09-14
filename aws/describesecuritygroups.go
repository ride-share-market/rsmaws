package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func DescribeSecurityGroups() {

	//	svc := ec2.New(nil)
	svc := ec2.New(&aws.Config{
		Region: aws.String("ap-southeast-1"),
	})

	params := &ec2.DescribeSecurityGroupsInput{}
	//	params := &ec2.DescribeVpcsInput{}

	resp, err := svc.DescribeSecurityGroups(params)
	//	resp, err := svc.DescribeVpcs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)

}
