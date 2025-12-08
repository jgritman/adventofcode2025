package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const DigitArraySize = 12

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
		var digitArray [DigitArraySize]rune
		line := scanner.Text()
		for i := 0; i < DigitArraySize; i++ {
			digitArray[i] = '0'
		}
		for i := 0; i < len(line); i++ {
			digitChanged := false
			arrayStart := DigitArraySize - (len(line) - i) 
			if arrayStart < 0 {
				arrayStart = 0	
			}
			currentRune := rune(line[i])
			for j := arrayStart; j < DigitArraySize; j++ {
				if digitChanged {
					digitArray[j] = '0'
				} else if currentRune > digitArray[j] {
					digitArray[j] = currentRune
					digitChanged = true
				}
			}
		}
		lineResult := 0
		for _, char := range digitArray {
			digit := int(char - '0')
			lineResult = lineResult * 10 + digit
		}
		// println(lineResult)
		total += lineResult
	}

  fmt.Println(total)	

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}
