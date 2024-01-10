package util

import "crypto/rand"

func RandomToken() ([]byte, error) {
	token := make([]byte, 16)
	_, err := rand.Read(token)
	if err != nil {
		return []byte{}, err
	}
	return token, nil
}
