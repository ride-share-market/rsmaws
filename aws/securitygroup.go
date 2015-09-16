package aws

import (
//	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/rudijs/rsmaws/trace"
)

type SecurityGroup struct {
	id           string
	name         string
	description  string
	vpcId        string
	ingressRules []securityGroupIngressRule
}

type securityGroupIngressRule struct {
	groupId     string
	protocol    string
	port        string
	sourceGroup string
	cidr        string
}

type AwsCreateSecurityGroupIface interface {
	Create(tracer trace.Tracer, svc *ec2.EC2, securityGroup *SecurityGroup) (string, error)
}

func NewSecurityGroup(name string, description string, vpcId string) *SecurityGroup {
	sg := &SecurityGroup{
		name:        name,
		description: description,
		vpcId:       vpcId,
	}

	sg.ingressRules = make([]securityGroupIngressRule, 5)

	return sg
}

func CreateSecurityGroups(awssg AwsCreateSecurityGroupIface, tracer trace.Tracer, svc *ec2.EC2, vpcId string) (map[string]*SecurityGroup, error) {
	
	securityGroupsList := []string{"NATSG", "DBServersSG",}
	
	// return value
	securityGroups := make(map[string]*SecurityGroup)

	for _, val := range securityGroupsList {
		
		natSg := NewSecurityGroup(val, val + " Security Group", vpcId)
		
		sgId, err := awssg.Create(tracer, svc, natSg)
		
		if err != nil {
			return nil, err
		}
		
		natSg.id = sgId
		
		securityGroups[val] = natSg
				
	}

	return securityGroups, nil

}

type AwsCreateSecurityGroup struct{}

func (this AwsCreateSecurityGroup) Create(tracer trace.Tracer, svc *ec2.EC2, securityGroup *SecurityGroup) (string, error) {

	params := &ec2.CreateSecurityGroupInput{
		Description: aws.String(securityGroup.description),
		GroupName:   aws.String(securityGroup.name),
		VpcId:       aws.String(securityGroup.vpcId),
	}

	tracer.Trace("Creating Security Group " + securityGroup.name + "...")
	resp, err := svc.CreateSecurityGroup(params)

	if err != nil {
		return "", err
	}

	tracer.Trace("Created Security Group " + securityGroup.name)

	return *resp.GroupId, nil

}
