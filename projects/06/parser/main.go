package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalf("File not provided as the second argument")
	}

	fileName := os.Args[1]

    file, err := os.Open(fileName)
    if err != nil {
    	log.Fatalf("Could not open file %q for reading", fileName)
	}

    parser := Parser{}
    data, err := parser.Parse(file)
    if err != nil {
    	log.Fatalf("Failed parsing! %s", err)
	}

	fmt.Print(string(data))
}
