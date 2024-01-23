package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: touch <file>")
		os.Exit(1)
	}

	filename := os.Args[1]

	err := createFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Printf("File '%s' touched successfully.\n", filename)
}

func createFile(filename string) error {
	f, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}
