/*
strife STatistical Recruiting Information For Engineers

Usage:
       strife <number> <path>
*/

package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/unlearned/strife"
)

func convNumType(s string) (uint16, error) {
	num, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return 0, err
	}

	if num <= 0 {
		return 0, errors.New("1st parameter MUST be over 0")
	}

	return uint16(num), nil
}

func main() {
	log.SetFlags(0)

	flag.Parse()
	flags := flag.Args()

	if len(flags) != 2 {
		log.Println(errors.New("Usage:\n       strife <number> <path>"))
		os.Exit(1)
	}

	jsonPath := flags[1]

	num, err := convNumType(flags[0])
	if err != nil && err != flag.ErrHelp {
		log.Println(err)
		os.Exit(1)
	}

	text, err := strife.PredictText(uint16(num), jsonPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Print(text)
}
