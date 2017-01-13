package main

import (
	"fmt"
	//	"github.com/aws/aws-sdk-go/aws"
	"flag"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

var ufile = flag.String("uploadfile", "", "File to upload to S3 bucket")
var dfile = flag.String("downloadfile", "", "File to download from S3 bucket")
var list = flag.Bool("list", false, "List Buckets on S3")

var s *session.Session
var svc *s3.S3

func main() {

	flag.Parse()

	s = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc = s3.New(s)

	if *list {
		printBuckets()
	}
}

func printBuckets() {
	result, err := svc.ListBuckets(nil)
	if err != nil {
		fmt.Println("Error fetching Instances from AWS", err)
		os.Exit(1)
	}
	fmt.Printf("Buckets\n")
	for _, b := range result.Buckets {
		fmt.Printf("%s\n", *b.Name)
	}
}
