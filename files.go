package go_utils

import (
	"context"
	"errors"
	"io"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
)

type FileService interface {
	Upload(ctx context.Context, file io.ReadCloser, key string, contentType string) (string, error)
	Sign(ctx context.Context, file string, eta time.Duration) (string, error)
}

type fileServiceImpl struct {
	client        *s3.Client
	bucketName    string
	presignClient *s3.PresignClient
}

// Sign implements FileService.
func (f *fileServiceImpl) Sign(ctx context.Context, file string, eta time.Duration) (string, error) {
	req, err := f.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(f.bucketName),
		Key:    aws.String(file),
	}, func(po *s3.PresignOptions) { po.Expires = eta })

	if err != nil {
		log.Printf("Couldn't get a presigned request to put %v:%v. Here's why: %v\n",
			f.bucketName, file, err)
	}

	return req.URL, err
}

// Upload implements FileService.
func (f *fileServiceImpl) Upload(ctx context.Context, file io.ReadCloser, key string, contentType string) (string, error) {
	_, err := f.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(f.bucketName),
		Key:    aws.String(key),
		Body:   file,
	})

	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode() == "EntityTooLarge" {
			log.Printf("Error while uploading object to %s. The object is too large.\n"+
				"To upload objects larger than 5GB, use the S3 console (160GB max)\n"+
				"or the multipart upload API (5TB max).", f.bucketName)
		} else {
			log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				key, f.bucketName, key, err)
		}
	} else {
		err = s3.NewObjectExistsWaiter(f.client).Wait(
			ctx, &s3.HeadObjectInput{Bucket: aws.String(f.bucketName), Key: aws.String(key)}, time.Minute)
		if err != nil {
			log.Printf("Failed attempt to wait for object %s to exist.\n", key)
		}
	}

	return key, err
}

func NewFileService(
	region string,
	baseEndpoint string,
	accessID string,
	accessSecret string,
	bucketName string,
) FileService {
	creds := credentials.NewStaticCredentialsProvider(accessID, accessSecret, "")

	ctx := context.Background()
	sdkConfig, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(region),
		config.WithBaseEndpoint(baseEndpoint),
		config.WithCredentialsProvider(creds),
	)

	if err != nil {
		panic(err)
	}

	s3Client := s3.NewFromConfig(sdkConfig)

	return &fileServiceImpl{
		client:        s3Client,
		bucketName:    bucketName,
		presignClient: s3.NewPresignClient(s3Client),
	}
}
