package file

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func S3Upload(bucket, region string) error {

	session, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return err
	}

	// Upload
	err = upload(session, bucket,"output.json")
	if err != nil {
		return err
	}

	return nil
}

func upload(session *session.Session, bucket, fileName string) error {
	log.Printf("Uploading to S3")

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(fileName),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		// ServerSideEncryption: aws.String("AES256"),
	})
	return err
}