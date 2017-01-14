package main

import (
	"fmt"
	//	"github.com/aws/aws-sdk-go/aws"
	"flag"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"path"
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

	b := flag.Args()
	if len(b) != 1 && !*list {
		fmt.Println("Must specify bucket name with --uploadfile or --downloadfile")
		usage()
	} else if len(b) > 0 {
		bucket = b[0]
	}

	if *list {
		if bucket == "" {
			printBuckets()
		} else {
			printBucketFiles()
		}
	}

	if *ufile != "" && *dfile != "" {
		fmt.Println("Cannot upload and download at the same time")
		usage()
	}

	if *ufile == "" && *dfile == "" && !*list {
		fmt.Println("Must specify --list, --uploadfile or --downloadfile")
		usage()
	}

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

func printBucketFiles() {
	result, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: &bucket})
	if err != nil {
		fmt.Println("Error fetching Objects from AWS", err)
		os.Exit(1)
	}
	fmt.Printf("Files of %s:\n", bucket)
	for _, b := range result.Contents {
		fmt.Printf("\t%s\n", *b.Key)
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

	biloader := s3manager.NewUploader(s)
	key := path.Base(*ufile)
	_, err = biloader.Upload(&s3manager.UploadInput{
		Bucket: &bucket,
		Key:    &key,
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
	key := *dfile
	file, err := os.Create(*dfile)
	if err != nil {
		fmt.Printf("Unable to create file %q\n", err)
		os.Exit(1)
	}

	fmt.Printf("Downloading...\n")
	biloader := s3manager.NewDownloader(s)
	_, err = biloader.Download(file, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		// Print the error and exit.
		fmt.Printf("Unable to download %q from %q, %v", *ufile, bucket, err)
		os.Exit(1)
	}

}
