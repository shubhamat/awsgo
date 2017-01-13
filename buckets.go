package main

import (
	"fmt"
	//	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func main() {
	s := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := s3.New(s)
	result, err := svc.ListBuckets(nil)
	if err != nil {
		fmt.Println("Error fetching Instances from AWS", err)
		os.Exit(1)
	} else {
		printBuckets(result)
	}
}

func printBuckets(o *s3.ListBucketsOutput) {
	fmt.Printf("Buckets\n")
	for _, b := range o.Buckets {
		fmt.Printf("%s\n", *b.Name)
	}
}
