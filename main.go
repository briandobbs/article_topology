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

	// Create an instance of TextBodyTopology
	counter := &service.TextBodyTopology{}

	// Parse the text to populate category counts
	counter.ParseText(byteStream)

	// Print the results
	fmt.Printf("Results: %v", counter.CharacterTypeSets)
}
