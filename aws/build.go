package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	
	"github.com/rudijs/rsmaws/trace"
	"os"
	
//	"fmt"
)

func BuildVpc(vpcRegion string) (map[string]string, error) {
	
	tracer := trace.New(os.Stdout)
	
	vpcData := make(map[string]string)

	svc := ec2.New(&aws.Config{
		Region: aws.String(vpcRegion),
	})

	// Create a new VPC
	awscvpc := new(AwsCreateVpc)
	vpcId, createVpcErr := CreateVpc(tracer, awscvpc, svc)
	if createVpcErr != nil {
		return nil, createVpcErr
	}
	vpcData["vpcId"] = vpcId
	
	// Get the Main Route Table ID
	routeTableId, describeRouteTablesErr := DescribeRouteTables(tracer, svc, vpcId)
	if describeRouteTablesErr != nil {
		return nil, describeRouteTablesErr
	}
	vpcData["routeTableId"] = routeTableId
	
	// Create Security Groups
	awssg := new(AwsCreateSecurityGroup)
	
	securityGroups, CreateSecurityGroupsErr := CreateSecurityGroups(awssg, tracer, svc, vpcData["vpcId"])
	
	if CreateSecurityGroupsErr != nil {
		return nil, CreateSecurityGroupsErr
	}
	
	vpcData["NATSG"] = securityGroups["NATSG"].id
	vpcData["DBServersSG"] = securityGroups["DBServersSG"].id


	// Create VPC Subnets
	
	// Elastic IP
	
	// Create private subnet NAT route to internet

	return vpcData, nil

}
