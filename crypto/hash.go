package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func HashSha256(value string) string {

	hasher := sha256.New()
	hasher.Write([]byte(value))
	hash := hasher.Sum(nil)

	return hex.EncodeToString(hash)
}

func HashMd5(value string) string {

	hasher := md5.New()
	hasher.Write([]byte(value))
	hash := hasher.Sum(nil)

	return hex.EncodeToString(hash)
}
