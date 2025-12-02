package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filePath := "../input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	position := 50
	zeroCount := 0
	
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0]
		clicks, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("Error parsing number: %v\n", err)
			return
		}
		direction := 1
		if dir == 'L' {
			direction = -1
		}
		rawPosition := (position + (clicks * direction)) % 100
		if rawPosition == 0 {
			zeroCount++
			position = 0
		}
		if rawPosition < 0 {
			position = 100 + rawPosition
		}
		if rawPosition > 0 {
			position = rawPosition 
		} 
	}

	fmt.Println(zeroCount)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}
