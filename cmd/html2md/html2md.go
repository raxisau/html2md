package main

import (
	"log"
	"os"

	"bitbucket.org/jackbooted/html2markdown/internal/pkg/html2md"
)

func main() {
	if err := html2md.Convert(os.Stdout, os.Stdin); err != nil {
		log.Fatal(err)
	}
}
