package aws

import (
	"bytes"
	"orderServiceCadenceWorker/v1/models"

	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"os"
)

type S3Client struct {
	s3     *s3.S3
	bucket string
	format string
}

func NewS3Client(accessKeyId string, secretAccessKey string, bucketRegion string, bucket string, format string) Client {
	dataLakeSession := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(bucketRegion),
		Credentials: credentials.NewStaticCredentials(accessKeyId, secretAccessKey, ""),
	}))

	client := s3.New(dataLakeSession)

	return &S3Client{
		s3:     client,
		bucket: bucket,
		format: format,
	}
}

func NewS3ClientWithConfigurations(configs []byte) (Client, error) {
	var s3Configurations models.S3ConfigurationDetails

	err := json.Unmarshal(configs, &s3Configurations)
	if err != nil {
		return nil, err
	}

	dataLakeSession := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(s3Configurations.S3BucketRegion),
		Credentials: credentials.NewStaticCredentials(s3Configurations.AccessKeyId, s3Configurations.SecretAccessKey, ""),
	}))

	client := s3.New(dataLakeSession)

	return &S3Client{
		s3:     client,
		bucket: s3Configurations.S3BucketName,
		format: s3Configurations.Format.FormatType,
	}, nil
}

func (c *S3Client) UploadFile(s3Key string, toUploadFilePath string) error {
	upFile, err := os.Open(toUploadFilePath)
	if err != nil {
		return err
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	fileSize := upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)

	_, err = upFile.Read(fileBuffer)
	if err != nil {
		return err
	}

	_, err = c.s3.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(c.bucket),
		Key:           aws.String(s3Key),
		Body:          bytes.NewReader(fileBuffer),
		ContentLength: aws.Int64(fileSize),
		ContentType:   aws.String(http.DetectContentType(fileBuffer)),
	})

	return err
}

func (c *S3Client) DeleteDirectoryByPrefix(prefix string) error {
	params := &s3.ListObjectsInput{
		Bucket: aws.String(c.bucket),
		Prefix: aws.String(prefix),
	}

	resp, err := c.s3.ListObjects(params)
	if err != nil {
		return err
	}

	for _, key := range resp.Contents {
		_, err = c.s3.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(c.bucket),
			Key:    key.Key,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *S3Client) DeleteS3ObjectByPath(path string) error {
	_, err := c.s3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    &path,
	})
	if err != nil {
		return err
	}

	return nil
}
