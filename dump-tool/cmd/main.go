package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	dir    string = "/tmp/haha"
	bucket string = "sn-dump-collector"
)

func main() {
	for range time.Tick(time.Second) {
		fmt.Print("start watch files\n")
		err := watch()
		if err != nil {
			fmt.Printf("watch files failed, %v\n", err)
		}
	}
}

func getNS() string {
	os.Getenv("NAMESPACE_NAME")
}

func getPodName() string {
	return os.Getenv("HOSTNAME")
}

func upload(fileName string) error {
	fmt.Printf("start to upload files %s\n", fileName)
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("upload canceled due to openning file, %v\n", err)
		return err
	}

	key := filepath.Join(time.Now().Format(time.RFC3339), getNS(), getPodName(), fileName)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			// If the SDK can determine the request or retry delay was canceled
			// by a context the CanceledErrorCode error code will be returned.
			fmt.Printf("upload canceled due to timeout, %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		}
		return err
	}

	fmt.Printf("successfully uploaded file to %s/%s\n", bucket, key)
	return nil
}

func watch() error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, fi := range files {
		fileName := filepath.Join(dir, fi.Name())
		err := upload(fileName)
		if err != nil {
			return err
		}
		err = os.Remove(fileName)
		if err != nil {
			return err
		}
	}
	return nil
}
