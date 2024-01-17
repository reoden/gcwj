package main

import (
	"fmt"
	"log"
	"os"
)

func cleanup() {
	if FilePtr != nil {
		fmt.Println("Closing the file")
		FilePtr.Close()
	}
}

func usage(prog string) {
	log.Printf("Usage: %s <infile>\n", prog)
	os.Exit(1)
}

func initVariable() {
	Line = 1
	Putback = '\n'
}

func main() {
	defer cleanup()

	if len(os.Args) != 2 {
		usage(os.Args[0])
	}

	initVariable()
	openfile(os.Args[1])
	scanFile()
}
