package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
)

func Decrypt(input string, key *rsa.PrivateKey) []byte {
	privateKey := x509.MarshalPKCS1PrivateKey(key)

	hash := sha256.Sum256(privateKey)
	aesKey := hash[:aes.BlockSize]

	block, err := aes.NewCipher(aesKey)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	inputBytes := []byte(input)

	iv := inputBytes[:aes.BlockSize]

	cipherInput := inputBytes[aes.BlockSize:]

	cipherText := make([]byte, len(cipherInput))

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherInput)

	return cipherText
}
