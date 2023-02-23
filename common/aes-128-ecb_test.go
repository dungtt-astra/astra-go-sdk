package common

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestECBEncrypt(t *testing.T) {
	key, _ := GenerateSecretKeyRandomString(32)
	keyByte, _ := DecodeBase64(key)

	data := "mammal initial effort joke public daring fish puppy risk famous cream occur else busy cable cruel vacant brick used patient choose object teach special"
	result, err := ECBEncrypt([]byte(data), keyByte)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	result1, err := ECBDecrypt(result, keyByte)

	if err != nil {
		panic(err)
	}

	fmt.Println(result1)
	assert.Equal(t, data, result1)
}

func TestGenerateKey(t *testing.T) {
	key, _ := GenerateSecretKeyRandomString(32)
	iv, _ := GenerateSecretKeyRandomString(16)

	fmt.Println(key, iv)
}
