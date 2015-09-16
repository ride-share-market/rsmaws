package aws

import (
//	"errors"
	"testing"
	"os"
	
	"github.com/aws/aws-sdk-go/service/ec2"
	
	"github.com/rudijs/rsmaws/trace"	
)

type Vpc struct {
	VpcId *string
}

type CreateVpcOutput struct {
	Vpc *Vpc
}

type createVpc struct{}

func (this createVpc) Create(tracer trace.Tracer, svc *ec2.EC2) (string, error) {

	vpcId := "vpc-abc123"

	vpc := &Vpc{
		VpcId: &vpcId,
	}

	vpcOutput := &CreateVpcOutput{
		Vpc: vpc,
	}
	
	return *vpcOutput.Vpc.VpcId, nil

}

func TestCreateVpc(t *testing.T) {
	
	tracer := trace.New(os.Stdout)

	svc := ec2.New(nil)
	
	vpc := new(createVpc)
	
	vpcId, _ := CreateVpc(tracer, vpc, svc)
	
	if vpcId != "vpc-abc123" {
		t.Errorf("vpcId != vpc-abc123 got %s", vpcId)
	}
	
}
