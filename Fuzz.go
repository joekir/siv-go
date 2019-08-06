package siv

import (
	"crypto/aes"
	"encoding/hex"
)

func Fuzz(data []byte) int {
	key, _ := hex.DecodeString("fffefdfcfbfaf9f8f7f6f5f4f3f2f1f0f0f1f2f3f4f5f6f7f8f9fafbfcfdfeff")
	aead, err := New(key, aes.NewCipher)
	if err != nil {
		return 0
	}

	// TODO fuzz first args too
	ciphertext := aead.Seal(nil, nil, data, data)
	data[0] ^= 1

	_, err = aead.Open(nil, nil, ciphertext, data)
	if err == nil {
		return 0
	}

	return 1
}
