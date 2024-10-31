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

	// Create an instance of CharCategoryCounter
	counter := &service.CharCategoryCounter{}

	// Parse the text to populate category counts
	counter.ParseText(byteStream)

	// Print the results
	fmt.Printf("Alphabet: %v\n", counter.Alphabet)
	fmt.Printf("Digits: %v\n", counter.Digit)
	fmt.Printf("Whitespace: %v\n", counter.Whitespace)
	fmt.Printf("Control: %v\n", counter.Control)
	fmt.Printf("Special: %v\n", counter.Special)
}
