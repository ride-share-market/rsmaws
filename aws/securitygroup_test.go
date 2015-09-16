package aws

import (
	"os"
	"testing"
	//    "errors"

	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/rudijs/rsmaws/trace"

)

type awsCreateSecurityGroup struct{}

func (this awsCreateSecurityGroup) Create(tracer trace.Tracer, svc *ec2.EC2, securityGroup *SecurityGroup) (string, error) {
	return "sg-abc123", nil
}

func TestCreateSecurityGroupsOK(t *testing.T) {

	tracer := trace.New(os.Stdout)

	svc := ec2.New(nil)

	vpcId := "vpc-abc123"

	awssg := new(awsCreateSecurityGroup)

	securityGroups, _ := CreateSecurityGroups(awssg, tracer, svc, vpcId)

	if securityGroups["NATSG"].name != "NATSG" {
		t.Errorf("Security Group name should be NATSG got %s", securityGroups["NATSG"].name)
	}

	if securityGroups["DBServersSG"].name != "DBServersSG" {
		t.Errorf("Security Group name should be DBServersSG got %s", securityGroups["DBServersSG"].name)
	}

}
