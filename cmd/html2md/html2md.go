package main

import (
	"log"
	"os"

	"github.com/raxisau/html2md/internal/pkg/html2md"
)

func main() {
	if err := html2md.Convert(os.Stdout, os.Stdin); err != nil {
		log.Fatal(err)
	}
}
