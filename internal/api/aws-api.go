package api

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (api *API) UploadQRCode(qrCode []byte, userId string) (key, objectURL string, err error) {
	tmpQrCode, err := os.CreateTemp("", "qr-*.png")
	if err != nil {
		return "", "", err
	}
	defer tmpQrCode.Close()
	defer os.Remove(tmpQrCode.Name())

	if _, err := tmpQrCode.Write(qrCode); err != nil {
		return "", "", err
	}

	tmpQrCode.Seek(0, 0)
	key = filepath.Base(tmpQrCode.Name())

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", "", err
	}

	s3Client := s3.NewFromConfig(cfg)

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(api.config.BUCKET_NAME),
		Key:    aws.String("qr-codes/" + userId + "/" + key),
		Body:   tmpQrCode,
	})
	if err != nil {
		return "", "", err
	}

	// Get the object URL
	objectURL = api.GetObjectURL(api.config.BUCKET_NAME, "qr-codes/"+userId+"/"+key)

	return key, objectURL, nil
}

// GetObjectURL retrieves the URL of the object stored in S3.
func (api *API) GetObjectURL(bucketName, key string) string {
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, key)
}