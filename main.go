package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	awsSession := session.New()
	s3Service := s3.New(awsSession)

	resp, err := s3Service.GetBucketLocation(&s3.GetBucketLocationInput{
		Bucket: aws.String("s3up-test"),
	})
	fmt.Println("RESP:")
	fmt.Println(resp)
	fmt.Println("ERR:")
	fmt.Println(err)
}
