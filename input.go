package main

import (
	"flag"
	"fmt"
)

func ParseInput() string {
	var input string

	flag.StringVar(&input, "name", "", "file name of the csv")
	flag.Parse()

	if len(input) == 0 {
		fmt.Println("Please provide a file name")

		return ""
	}

	return input
}
