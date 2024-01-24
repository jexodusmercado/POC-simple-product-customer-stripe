package helper

import "math/rand"

// generate a random numbers with a prefix of "KC-" and a length of 8
func GenerateOrderID() string {
	return "KC-" + GenerateRandomString(8)
}

// generate a random string with a length of n
func GenerateRandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[GenerateRandomInt(len(letters))]
	}
	return string(b)
}

// generate a random int with a max of n
func GenerateRandomInt(n int) int {
	return rand.Intn(n)
}
