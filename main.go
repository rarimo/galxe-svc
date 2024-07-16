package main

import (
	"os"

	"github.com/rarimo/galxe-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
