package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"bitbucket.org/jackbooted/html2markdown/internal/pkg/html2md"
)

var (
	guesslang = flag.String("g", "", "guesslang")
	option    *html2md.Option
)

func guesslanger(code string) (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", *guesslang)
	} else {
		cmd = exec.Command("sh", "-c", *guesslang)
	}
	cmd.Stdin = strings.NewReader(code)
	b, err := cmd.CombinedOutput()
	return strings.ToLower(strings.TrimSpace(string(b))), err
}

func main() {
	flag.Parse()
	if *guesslang != "" {
		option = &html2md.Option{GuessLang: guesslanger}
	}
	option := &html2md.Option{GuessLang: guesslanger}
	if err := html2md.Convert(os.Stdout, os.Stdin, option); err != nil {
		log.Fatal(err)
	}
}
