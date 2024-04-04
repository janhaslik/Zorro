package cli

import (
	"crypto/rsa"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"zoro/crypto"
	"zoro/db"
)

func ProcessED(operation string, args []string) {
	if len(args) < 3 {
		log.Fatal("Usage: zoro", operation, "<rsa_key> <input_file> <output_file>")
	}

	action := args[0]

	if action == "-h" || action == "--help" || action == "help" {
		switch operation {
		case "encrypt":
			helpEncrypt()
		case "decrypt":
			helpDecrypt()
		default:
			Help()
		}
		return
	}

	rsaKeyInput := action
	input := args[1]
	outputFile := args[2]

	var rsaKey *rsa.PrivateKey
	var err error

	// Check if the RSA key input starts with '--db-'
	if strings.HasPrefix(rsaKeyInput, "--db-") {
		id, _ := strconv.Atoi(rsaKeyInput[5:])

		rsaKeyStr, err := db.GetRSAKey(id)
		rsaKey = crypto.ParsePrivateKey(rsaKeyStr)

		if err != nil {
			log.Fatal("Error retrieving RSA key from database:", err)
		}
	} else {
		rsaKey = crypto.ReadKeyFromFile(rsaKeyInput)
	}

	if strings.Contains(input, ".txt") {
		fileData, err := os.ReadFile(input)
		if err != nil {
			log.Fatal("Error reading input file:", err)
		}

		input = string(fileData)
	}

	var result []byte
	switch operation {
	case "encrypt":
		result = crypto.Encrypt(input, rsaKey)
	case "decrypt":
		result = crypto.Decrypt(input, rsaKey)
	}

	err = os.WriteFile(outputFile, result, 0644)
	if err != nil {
		log.Fatal("Error writing to output file:", err)
	}
}

func ProcessRSAKey(args []string) {
	if len(args) < 1 {
		log.Fatal("Usage: go run main.go rsa-key <action> [args]")
	}

	action := args[0]
	switch action {
	case "generate":
		processGenerateRSAKey(args[1:])
	case "list":
		processListDb()
	case "delete":
		processDeleteRSAKey(args[1:])

	default:
		log.Fatal("Invalid action for rsa-key command. Please specify 'generate' or 'list'.")
	}
}

func processGenerateRSAKey(args []string) {
	if len(args) < 1 {
		helpRSAKeyGenerate()
	}

	action := args[0]

	if action == "-h" || action == "--help" || action == "help" {
		helpRSAKeyGenerate()
		return
	}

	bitLength, err := strconv.Atoi(action)
	if err != nil {
		log.Fatal("Invalid bit length:", err)
	}

	pemFile := "private_key.pem"
	if len(args) > 1 {
		pemFile = args[1]
	}

	key := crypto.GenerateRSAKey(bitLength, pemFile)

	if len(args) > 2 && args[2] == "--save" {
		keyName := args[3]
		crypto.SaveKeyToDb(key, keyName)
	}
}

func processDeleteRSAKey(args []string) {
	if len(args) > 0 {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return
		}
		db.DeleteRSAKey(id)
	}
}

func ProcessHash(args []string) {
	if len(args) > 0 {
		method := args[0]
		value := args[1]
		var hash string
		switch method {
		case "sha256":
			hash = crypto.HashSha256(value)
		case "md5":
			hash = crypto.HashMd5(value)
		case "list":
			helpHash()
		default:
			helpHash()
			return
		}
		fmt.Println(hash)
	}
}
