package main

import (
	"flag"
	"fmt"

	"github.com/rudijs/rsmaws/aws"
)

func main() {

	//	type Vpc struct {
	//	CidrBlock *string
	//	}
	//
	//	type CreateVpcOutput struct {
	//		Vpc *Vpc
	//	}
	//
	//	val := "1.1.1.1"
	//
	//	a := &Vpc{
	//		CidrBlock: &val,
	//	}
	//
	//	println(*a.CidrBlock)
	//
	//	b := &CreateVpcOutput{
	//		Vpc: a,
	//	}
	//
	//	c := *b.Vpc.CidrBlock
	//
	//	println(c)
	//
	//
	//	return

	hostedZonePtr := flag.Bool("hostedZone", false, "Get AWS Hosted Zone")
	hostedZoneIdPtr := flag.String("hostedZoneId", "", "AWS Hosted Zone ID")

	vpcCreatePtr := flag.Bool("vpcCreate", false, "Create a new VPC environment")
	vpcRegionPtr := flag.String("vpcRegion", "", "AWS VPC Region")

	flag.Parse()

	var cliMode string

	if *hostedZonePtr == true {
		cliMode = "hostedZone"
	}

	if *vpcCreatePtr == true {
		cliMode = "vpcCreate"
	}

	if cliMode == "" {
		flag.Usage()
		return
	}

	switch cliMode {

	case "hostedZone":

		//		fmt.Println("hostedZone")

		hostedZoneId := *hostedZoneIdPtr

		if hostedZoneId == "" {
			panic("--hostedZoneId is required")
		}

		awshz := new(aws.AwsHostedZone)

		hz, err := aws.GetHostedZone(awshz, hostedZoneId)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(hz)

	case "vpcCreate":

		if *vpcRegionPtr == "" {
			panic("--vpcRegion is required")
		}

		res, err := aws.BuildVpc(*vpcRegionPtr)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(res)

	}

}
