package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	
	"github.com/rudijs/rsmaws/trace"
)

type CreateVpcIface interface {
	Create(tracer trace.Tracer, svc *ec2.EC2) (string, error)
}

func CreateVpc(tracer trace.Tracer, vpc CreateVpcIface, svc *ec2.EC2) (string, error) {
	return vpc.Create(tracer, svc)
}

type AwsCreateVpc struct{}

func (this AwsCreateVpc) Create(tracer trace.Tracer, svc *ec2.EC2) (string, error) {
	
	params := &ec2.CreateVpcInput{
		CidrBlock: aws.String("10.0.0.0/16"),
	}

	tracer.Trace("Creating AWS VPC...")
	resp, err := svc.CreateVpc(params)

	if err != nil {
		return "", err
	}

	// Pretty-print the response data.
	//	fmt.Println(resp)
	
	vpcId := *resp.Vpc.VpcId
	tracer.Trace("Created AWS VPC: " + vpcId)

	tracer.Trace("Modifying AWS VPC Attributes...")
	errModify := this.ModifyVpcAttribute(svc, vpcId)

	if errModify != nil {
		return "", errModify
	}
	tracer.Trace("AWS VPC Attributes Modified.")

	return vpcId, nil

}

func (this AwsCreateVpc) ModifyVpcAttribute(svc *ec2.EC2, vpcId string) (error) {

	params := &ec2.ModifyVpcAttributeInput{
		VpcId: aws.String(vpcId),
		EnableDnsHostnames: &ec2.AttributeBooleanValue{
			Value: aws.Bool(true),
		},
	}

	_, err := svc.ModifyVpcAttribute(params)

	if err != nil {
		return err
	}
	
	return nil

}
