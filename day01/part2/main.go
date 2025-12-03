package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
		fmt.Println(line)
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
		rawClicks := (position + (clicks * direction)) 
		zeroCount += int(math.Abs(float64(rawClicks / 100)))	
		rawPosition := rawClicks % 100
		if rawPosition == 0 && direction == -1 {
			zeroCount++
		}
		if rawPosition < 0 {
			if position != 0 {
				zeroCount++
			}
			position = 100 + rawPosition
		} else {
			position = rawPosition 
		}
		fmt.Println(position, zeroCount)
	}

	fmt.Println(zeroCount)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}
