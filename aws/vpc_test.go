package aws

import (
//	"errors"
	"testing"
	"fmt"
)

type Vpc struct {
	VpcId *string
}

type CreateVpcOutput struct {
	Vpc *Vpc
}

type createVpc struct{}

func (this createVpc) Create(vpcRegion string) (string, error) {

	vpcId := "vpc-b9258cdc"

	vpc := &Vpc{
		VpcId: &vpcId,
	}

	vpcOutput := &CreateVpcOutput{
		Vpc: vpc,
	}
	
	return *vpcOutput.Vpc.VpcId, nil

}

func TestCreateVpc(t *testing.T) {
	
	vpc := new(createVpc)
	
	vpcRegion := "ap-southeast-1"
	
	vpcId, _ := CreateVpc(vpc, vpcRegion)
	
	fmt.Println(vpcId)

}
