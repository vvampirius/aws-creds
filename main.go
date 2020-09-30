package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const VERSION  = 0.1

func helpText() {
	fmt.Println("https://github.com/vvampirius/aws-creds\n")
	fmt.Println("Get password from keychain and return it in aws-ready format\n")
	fmt.Println("USAGE:")
	fmt.Printf(" %s keychain <keychain> <name>\n\n", os.Args[0])
	flag.PrintDefaults()
}

func wrongUsage() {
	os.Stdout = os.Stderr
	helpText()
	os.Exit(1)
}

func main() {
	help := flag.Bool("h", false, "print this help")
	ver := flag.Bool("v", false, "Show version")
	flag.Parse()

	if *help {
		helpText()
		os.Exit(0)
	}

	if *ver {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	log.SetFlags(log.Lshortfile)
	log.SetOutput(os.Stderr)

	switch flag.Arg(0) {
	case `keychain`:
		if len(flag.Args()) != 3 { wrongUsage() }
		keychainPath := flag.Arg(1)
		name := flag.Arg(2)
		password, err := GetPasswordFromKeychain(name, keychainPath)
		if err != nil { os.Exit(2) }

		credentials := NewCredentials(name, string(password))
		if err := credentials.Fprint(os.Stdout); err != nil {
			os.Exit(3)
		}
	default:
		wrongUsage()
	}
}
