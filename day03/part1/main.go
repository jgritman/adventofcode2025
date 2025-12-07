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

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		firstChar := '0'
		secondChar := '0'
		for i, char := range line {
			if char > firstChar && len(line) - 1 != i {
				firstChar = char
				secondChar = '0'
			} else if char > secondChar {
				secondChar = char
			}
		}
		lineResult := int(firstChar-'0')*10 + int(secondChar-'0')
		// println(lineResult)
		total += lineResult
	}

  fmt.Println(total)	

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}
