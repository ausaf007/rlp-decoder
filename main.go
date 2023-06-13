package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"rlp-decoder/rlp"
)

const DISCLAIMER = "\nDisclaimer(For Windows users only):" +
	"\nPlease note Windows CMD or Powershell might have a input limit of 255 characters." +
	"\nIf you are on windows, consider using WSL or Git Bash Emulator to run this program."

// loadFlags that specifies the verbosity level of the program based on user specified flag
func loadFlags() {
	isVerbose := flag.Bool("verbose", false, "Specifies verbosity of logs. True means Info Level. "+
		"False means Warn Level")
	flag.Parse()
	if *isVerbose {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}

// printError prints the error and asks user to try again
func printError(err error) {
	fmt.Println("\nError Encountered:", err,
		"\nPlease try again.")
}

// driver code
func main() {

	loadFlags()
	fmt.Println(DISCLAIMER)

	// This keeps iterating the loop until a valid input string is entered,
	// and an RLP encoded string is decoded successfully
	for true {

		// Assuming the string entered is in hex format, with no 0x prefix
		fmt.Println("\nEnter string to be decoded: ")

		//input string that is to be decoded
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			printError(err)
			continue
		}

		decodedStr, err := rlp.Decode(input)
		if err != nil {
			fmt.Println("Decoded Values so far:\n", decodedStr)
			printError(err)
			continue
		}

		fmt.Println(decodedStr)
		break
	}
	fmt.Println("Decoding Successful!")
}
