package main

import (
	"fmt"
	//	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"os"
	"time"
)

func main() {
	s := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	esvc := ec2.New(s)
	result, err := esvc.DescribeInstances(nil)
	//re := regexp.MustCompile("LaunchTime*,")
	if err != nil {
		fmt.Println("Error fetching Instances from AWS", err)
		os.Exit(1)
	} else {
		printInstanceUpTime(result)
	}
}

func printInstanceUpTime(desc *ec2.DescribeInstancesOutput) {
	fmt.Printf("PublicIpAddress\t\tUpTime\n")
	for _, r := range desc.Reservations {
		for _, i := range r.Instances {
			if *i.State.Name == "running" {
				fmt.Printf("%s\t\t%s\n", *i.PublicIpAddress, time.Since(*i.LaunchTime))
			}
		}

	}

}
