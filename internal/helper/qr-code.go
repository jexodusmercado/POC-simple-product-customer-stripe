package helper

import (
	"encoding/json"
	"fmt"

	"github.com/skip2/go-qrcode"
)

type QRCodeDetails struct {
	UserID string
	// insert other fields here
}

func GenerateQRCode(qrCodeDetails QRCodeDetails) ([]byte, error) {
	json, err := json.Marshal(qrCodeDetails)
	if err != nil {
		fmt.Println("Error in GenerateQRCode: ", err)
		return nil, err
	}

	// generate QR code
	qr, err := qrcode.Encode(string(json), qrcode.Medium, 256)
	if err != nil {
		fmt.Println("Error in GenerateQRCode: ", err)
		return nil, err
	}

	return qr, nil
}
