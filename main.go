package main

import (
	"article_topology/service"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("README.md")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer f.Close() // Ensure the file is closed when we're done

	// Sample byte stream
	//byteStream := []byte{'a', '5', ' ', '#', 'Z', '9', '\n', '%'}
	byteStream, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, b := range byteStream {
		switch {
		case service.IsAlphabetChar(b):
			fmt.Printf("'%c' is an alphabet character\n", b)
		case service.IsDigit(b):
			fmt.Printf("'%c' is a digit\n", b)
		case service.IsControlChar(b):
			fmt.Printf("'%c' is a control character\n", b)
		case service.IsWhitespace(b):
			fmt.Printf("'%c' is a whitespace character\n", b)
		case service.IsSpecialChar(b):
			fmt.Printf("'%c' is a special character\n", b)
		default:
			fmt.Printf("'%c' is unclassified\n", b)
		}
	}
}
