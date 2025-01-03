package main

import (
	"os"

	"github.com/cifra-city/entities-storage/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
