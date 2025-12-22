package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "../input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1 // Initialize a line number counter

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line %d: %s\n", lineNumber, line)
		lineNumber++ // Increment the line number
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}