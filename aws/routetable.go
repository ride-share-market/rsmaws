package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	
	"github.com/rudijs/rsmaws/trace"
)

func DescribeRouteTables(tracer trace.Tracer, svc *ec2.EC2, vpcId string) (string, error) {

	params := &ec2.DescribeRouteTablesInput{
		Filters: []*ec2.Filter{
			{ // Required
				Name: aws.String("vpc-id"),
				Values: []*string{
					aws.String(vpcId),
				},
			},
		},
	}
	
	tracer.Trace("Getting Main Route Table ID...")
	resp, err := svc.DescribeRouteTables(params)

	if err != nil {
		return "", err
	}
	
	tracer.Trace("Main Route Table ID " + *resp.RouteTables[0].RouteTableId)

	return *resp.RouteTables[0].RouteTableId, nil

}
