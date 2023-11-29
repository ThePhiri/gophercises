package main

import (
	"flag"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")

	flag.Parse()

	file, err := os.Open(*csvFilename)

}
