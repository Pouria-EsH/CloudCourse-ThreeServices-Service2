package storage

import (
	"bytes"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ArvanCloudS3 struct {
	s3client      *s3.S3
	bucket        string
	bucketAddress string
}

func NewArvanCloudS3(bucketName string, region string, endpoint string, accessKeyID string, secretAccessKey string) (*ArvanCloudS3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(region),
		Endpoint:         aws.String(endpoint),
		Credentials:      credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		S3ForcePathStyle: aws.Bool(true), // Required for ArvanCloud S3 compatibility
	})
	if err != nil {
		return nil, fmt.Errorf("couldn't create new session: %w", err)
	}

	return &ArvanCloudS3{
		s3client:      s3.New(sess),
		bucket:        bucketName,
		bucketAddress: fmt.Sprintf("https://%s.s3.%s.arvanstorage.ir", bucketName, region),
	}, nil
}

func (p ArvanCloudS3) Download(key string) (*bytes.Buffer, error) {
	obj, err := p.s3client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(p.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	var buffer []byte
	buf := bytes.NewBuffer(buffer)
	_, err = io.Copy(buf, obj.Body)
	if err != nil {
		fmt.Println("cant copy file")
	}
	return buf, nil
}
