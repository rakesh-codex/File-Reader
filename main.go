package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Check if a file path is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file-path>")
		fmt.Println("Example: go run main.go sample.txt")
		os.Exit(1)
	}

	// Get the file path from the command-line argument
	filePath := os.Args[1]

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Create a scanner to read the file line-by-line
	scanner := bufio.NewScanner(file)

	// Read and display the file contents
	lineNumber := 1
	for scanner.Scan() {
		fmt.Printf("%d: %s\n", lineNumber, scanner.Text())
		lineNumber++
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// If the file is empty
	if lineNumber == 1 {
		fmt.Println("File is empty.")
	}
}
