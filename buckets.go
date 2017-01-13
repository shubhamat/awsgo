package main

import (
	"fmt"
	//	"github.com/aws/aws-sdk-go/aws"
	"flag"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3manager"
	"os"
)

var ufile = flag.String("uploadfile", "", "File to upload to S3 bucket")
var dfile = flag.String("downloadfile", "", "File to download from S3 bucket")
var list = flag.Bool("list", false, "List Buckets on S3")

var s *session.Session
var svc *s3.S3
var bucket string

func main() {

	flag.Parse()

	s = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc = s3.New(s)

	if *list {
		printBuckets()
	}

	if *ufile != "" && *dfile != "" {
		fmt.Println("Cannot upload and download at the same time")
		usage()
	}

	if *ufile == "" && *dfile == "" {
		fmt.Println("Must specify --uploadfile or --downloadfile")
		usage()
	}

	b := flag.Args()
	if len(b) != 1 {
		fmt.Println("Must specify bucket name with --uploadfile or --downloadfile")
		usage()
	}

	bucket = b[0]

	if *ufile != "" {
		upload()
	} else {
		download()
	}

}

func usage() {
	fmt.Println("Usage: bucket [--list] | [ --uploadfile=<filename> | --downloadfile=<filename> ] bucketname")
	os.Exit(1)
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
	os.Exit(0)
}

func upload() {
	file, err := os.Open(*ufile)
	if err != nil {
		fmt.Printf("Unable to open file %q\n", err)
		os.Exit(1)
	}
	defer file.Close()

	uploader := s3manager.NewUploader(s)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(*ufile),
		Body:   file,
	})
	if err != nil {
		// Print the error and exit.
		fmt.Printf("Unable to upload %q to %q, %v", *ufile, bucket, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", *ufile, bucket)
}

func download() {

}
