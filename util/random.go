package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
)

func GenerateRandomToken(length int) (string, error) {
	token := make([]byte, length)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}

func GenerateRandomPasscode(length int) (string, error) {
	bi, err := rand.Int(
		rand.Reader,
		big.NewInt(int64(math.Pow(10, float64(length)))),
	)
	if err != nil {
		return "", err
	}
	passcode := fmt.Sprintf("%0*d", length, bi)
	return passcode, nil
}
