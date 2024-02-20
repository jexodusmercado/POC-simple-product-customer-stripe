package api

import (
	"context"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (api *API) UploadQRCode(qrCode []byte, userId string) (string, error) {

	tmpQrCode, err := os.CreateTemp("", "qr-*.png")
	if err != nil {
		return "", err
	}
	defer tmpQrCode.Close()
	defer os.Remove(tmpQrCode.Name())

	if _, err := tmpQrCode.Write(qrCode); err != nil {
		return "", err
	}

	tmpQrCode.Seek(0, 0)
	key := filepath.Base(tmpQrCode.Name())

	_, err = api.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(api.config.BUCKET_NAME),
		Key:    aws.String("qr-codes"+ "/" + userId + "/" + key),
		Body:   tmpQrCode,
	})

	if err != nil {
		return "", err
	}

	return key, nil
}
