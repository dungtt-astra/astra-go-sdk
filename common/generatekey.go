package common

import (
	"crypto/rand"
)

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateSecretKeyRandomString(n int) (string, error) {
	key, err := generateRandomBytes(n)
	if err != nil {
		return "", err
	}

	return EncodeBase64(key), nil
}
