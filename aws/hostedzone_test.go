package aws

import (
	"testing"
	"errors"
)

type awsHostedZone struct {}

func (this awsHostedZone) Get(s string) (string, error) {
	if s == "error" {
		return "", errors.New("Something Blew Up!")
	} else {
		return `{"a":"b"}`, nil
	}
}

func TestGetHostedZoneSuccess(t *testing.T) {

	awshz := &awsHostedZone{}

	in := "abc123"
	out := `{"a":"b"}`

	hz, _ := GetHostedZone(awshz, in)

	if hz != out {
		t.Errorf("%q != %q", hz, out)
	}

}

func TestGetHostedZoneFail(t *testing.T) {

	awshz := &awsHostedZone{}

	in := "error"

	_, err := GetHostedZone(awshz, in)

	if err == nil {
		t.Errorf("Error Expected")
	}

}
