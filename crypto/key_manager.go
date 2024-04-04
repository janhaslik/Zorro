package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"zoro/db"
)

func GenerateRSAKey(bits int, pemfile string) *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		fmt.Println("Error generating private key...")
		return nil
	}
	SaveKeyToPemFile(privateKey, pemfile)
	return privateKey
}

func SaveKeyToPemFile(key *rsa.PrivateKey, pemfile string) {
	file, err := os.Create(pemfile)

	if err != nil {
		fmt.Println("Error creating pem file...")
	}

	block := &pem.Block{
		Type:    "Message",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(key),
	}
	err = pem.Encode(file, block)

	if err != nil {
		fmt.Println("Error saving key to pem file...")
	}
}

func SaveKeyToDb(key *rsa.PrivateKey, name string) {
	block := &pem.Block{
		Type:    "Message",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(key),
	}

	pemBytes := pem.EncodeToMemory(block)
	pemStr := string(pemBytes)
	db.SaveRSAKey(name, pemStr)
}

func ReadKeyFromFile(pemfile string) *rsa.PrivateKey {
	file, err := os.ReadFile(pemfile)

	if err != nil {
		fmt.Println("Error creating pem file...")
	}

	keyBlock, _ := pem.Decode(file)

	key, _ := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)

	return key
}

func ParsePrivateKey(key string) *rsa.PrivateKey {
	keyBytes := []byte(key)
	block, _ := pem.Decode(keyBytes)

	if block == nil {
		fmt.Println("Error decoding PEM block")
		return nil
	}

	//parse key
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		fmt.Println("Error parse privatekey: ", err)
		return nil
	}

	return privateKey
}
