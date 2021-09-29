package cmd

import (
	"log"
	"os"
)

func ExitWithError(message string, err error) {
	log.Printf("%s: %s", message, err.Error())
	os.Exit(1)
}
