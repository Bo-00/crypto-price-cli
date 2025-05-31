package main

import (
	"os"

	"github.com/yourusername/crypto-price-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
