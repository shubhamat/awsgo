package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	s := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	ec2Svc := ec2.New(s)
	result, err := ec2Svc.DescribeInstanceStatus(nil)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success", result)
	}
}
