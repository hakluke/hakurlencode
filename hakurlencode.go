package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	decodeBool := flag.Bool("d", false, "URL decode, instead of encoding")
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)

	if *decodeBool {
		decode(*s)
	} else {
		encode(*s)
	}

}

func decode(s bufio.Scanner) {
	for s.Scan() {
		value, err := url.QueryUnescape(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(value)
	}
}

func encode(s bufio.Scanner) {
	for s.Scan() {
		fmt.Println(url.QueryEscape(s.Text()))
	}
}
