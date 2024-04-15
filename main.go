package main

import (
	"flag"
	"log"
	"zorro/cli"
)

func main() {
	flag.Usage = func() {
		cli.Help()
	}

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}

	command := args[0]

	switch command {
	case "encrypt", "decrypt":
		cli.ProcessED(command, args[1:])
	case "hash":
		cli.ProcessHash(args[1:])
	case "rsa-key":
		cli.ProcessRSAKey(args[1:])
	case "zoro-lost", "lorenor-zoro-moment", "zoro-knows-the-way", "help", "--help", "-h":
		cli.Help()
	default:
		log.Fatal("Invalid command. Please specify 'encrypt', 'decrypt', or 'rsa-key'.")
	}
}
