package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"io"
)

func Encrypt(input string, key *rsa.PrivateKey) []byte {
	privateKey := x509.MarshalPKCS1PrivateKey(key)

	hash := sha256.Sum256(privateKey)
	aesKey := hash[:aes.BlockSize]

	block, err := aes.NewCipher(aesKey)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	iv := make([]byte, aes.BlockSize)
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil
	}

	cipherInput := []byte(input)
	cipherInput = append(cipherInput, byte(aes.BlockSize-len(cipherInput)%aes.BlockSize))

	cipherText := make([]byte, aes.BlockSize+len(cipherInput))

	copy(cipherText[:aes.BlockSize], iv)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], cipherInput)

	return cipherText
}
