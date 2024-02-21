package helper

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
	"encoding/base64"
)

func Encrypt(data []byte, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    blockSize := block.BlockSize()
    plaintext := pad(data, blockSize)

    mode := cipher.NewCBCEncrypter(block, iv)
    ciphertext := make([]byte, len(plaintext))
    mode.CryptBlocks(ciphertext, plaintext)

    // Convert byte slice to string
    encryptedString := base64.StdEncoding.EncodeToString(ciphertext)

    return encryptedString, nil
}

func Decrypt(data string, key []byte, iv []byte) (string, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    decodedData, err := base64.StdEncoding.DecodeString(data)
    if err != nil {
        return "", err
    }

    mode := cipher.NewCBCDecrypter(block, iv)
    decrypted := make([]byte, len(decodedData))
    mode.CryptBlocks(decrypted, decodedData)

    decrypted = unpad(decrypted)

    return string(decrypted), nil
}

func pad(data []byte, blockSize int) []byte {
    padding := blockSize - (len(data) % blockSize)
    padText := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(data, padText...)
}

func unpad(data []byte) []byte {
    length := len(data)
    padding := int(data[length-1])
    return data[:length-padding]
}