package aws

import (
	"errors"
	"encoding/json"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

type AwsHostedZoneIface interface {
	Get(s string) (string, error)
}

func GetHostedZone(awshz AwsHostedZoneIface, hostedZoneId string) (string, error) {

	resp, err := awshz.Get(hostedZoneId)

	if err != nil {
		return "", err
	}

	return resp, nil

}

type AwsHostedZone struct{}

func (this AwsHostedZone) Get(s string) (string, error) {

	svc := route53.New(nil)

	params := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(s), // Required
	}
	
	resp, err1 := svc.ListResourceRecordSets(params)

	if err1 != nil {
		return "", errors.New(err1.Error())
	}

	jsonData, err2 := json.Marshal(resp)
	
	if err2 != nil {
		return "", errors.New(err2.Error())
	}

	return string(jsonData), nil

}