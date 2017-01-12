package main

import (
	"fmt"
//	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
  "time"
)

func main() {
	s := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	ec2Svc := ec2.New(s)
  result, err := ec2Svc.DescribeInstances(nil)
  //re := regexp.MustCompile("LaunchTime*,")
	if err != nil {
		fmt.Println("Error", err)
	} else {
    fmt.Println(time.Since(*result.Reservations[0].Instances[0].LaunchTime))
	}
}
