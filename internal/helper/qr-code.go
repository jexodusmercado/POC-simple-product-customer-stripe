package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"strconv"
	"strings"

	"github.com/skip2/go-qrcode"
)

type QRCodeDetails struct {
	UserID            string
	ProductID         string
	TransactionID     string
	TransactionType   string
	UserName          string
	UserEmail         string
	IsUserEarlyAccess  bool
	ProductName       string
	Description       string
	Package           string
	BasePrice         string
	PriceWithDiscount string
	Date              string
}

func GenerateQRCode(encryptedData string) ([]byte, error) {
	// generate QR code
	qr, err := qrcode.New(string(encryptedData), qrcode.Medium)
	if err != nil {
		fmt.Println("Error in GenerateQRCode: ", err)
		return nil, err
	}

	qr.BackgroundColor = color.RGBA{255, 255, 255, 255} // White background
	foregroundColor, err := hexToRGBA("#A554C4")

	if err != nil {
		fmt.Println("Error in GenerateQRCode: ", err)
		return nil, err
	}
	qr.ForegroundColor = foregroundColor

	// Encode QR code to PNG byte slice
	png, err := qr.PNG(200) // Size of 256x256 pixels
	if err != nil {
		fmt.Println("Error in GenerateQRCode: ", err)
		return nil, err
	}

	return png, nil
}

func hexToRGBA(hex string) (color.RGBA, error) {
	// Remove '#' from hex code
	hex = strings.TrimPrefix(hex, "#")

	// Parse the hex code to get RGB values
	rgb, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return color.RGBA{}, err
	}

	// Extract individual RGB components
	r := uint8((rgb >> 16) & 0xFF)
	g := uint8((rgb >> 8) & 0xFF)
	b := uint8(rgb & 0xFF)

	// Create RGBA color
	rgba := color.RGBA{r, g, b, 255}
	return rgba, nil
}

func addImageToQRCode(qrCodeImage []byte, smallImage []byte) ([]byte, error) {
	// Decode QR code image
	qrImage, _, err := image.Decode(bytes.NewReader(qrCodeImage))
	if err != nil {
		return nil, err
	}

	// Decode small image
	smallImg, _, err := image.Decode(bytes.NewReader(smallImage))
	if err != nil {
		return nil, err
	}

	// Calculate position to center small image on QR code
	qrWidth := qrImage.Bounds().Dx()
	qrHeight := qrImage.Bounds().Dy()
	smallWidth := smallImg.Bounds().Dx()
	smallHeight := smallImg.Bounds().Dy()
	posX := (qrWidth - smallWidth) / 2
	posY := (qrHeight - smallHeight) / 2

	// Create RGBA image
	rgba := image.NewRGBA(qrImage.Bounds())

	// Draw QR code onto RGBA image
	draw.Draw(rgba, qrImage.Bounds(), qrImage, image.Point{}, draw.Src)

	// Draw small image onto RGBA image at calculated position
	draw.Draw(rgba, smallImg.Bounds().Add(image.Point{posX, posY}), smallImg, image.Point{}, draw.Over)

	// Encode RGBA image to PNG byte slice
	var buf bytes.Buffer
	if err := png.Encode(&buf, rgba); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func GenerateQRCodeWithEncryptedData(qrCodeDetails QRCodeDetails, key, iv []byte, smallImage []byte) ([]byte, error) {
	// Encrypt QR code details
	jsonBytes, err := json.Marshal(qrCodeDetails)
	if err != nil {
		return nil, err
	}

	encryptedData, err := Encrypt(jsonBytes, key, iv)
	if err != nil {
		return nil, err
	}
	fmt.Println("encryptedData: ", encryptedData)

	descryptedData, err := Decrypt(encryptedData, key, iv)
	if err != nil {
		return nil, err
	}
	fmt.Println("descryptedData: ", descryptedData)

	qrCodeImage, err := GenerateQRCode(encryptedData)
	if err != nil {
		return nil, err
	}

	compositeImage, err := addImageToQRCode(qrCodeImage, smallImage)
	if err != nil {
		return nil, err
	}

	return compositeImage, nil
}
