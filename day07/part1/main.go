package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filePath := "../input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return
	}
	
	first := scanner.Text()
	startPos := strings.IndexRune(first, 'S')
	positions := map[int]struct{} { startPos: {} }
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		nextPositions := make(map[int]struct{})
		for pos := range positions {
			if line[pos] == '^' {
				total++
				leftSplit := pos - 1
				if leftSplit >= 0 {
					nextPositions[leftSplit] = struct{}{}
				} 
				rightSplit := pos + 1
				if rightSplit < len(line) {
					nextPositions[rightSplit] = struct{}{}
				}
			} else {
				nextPositions[pos] = struct{}{}
			} 
		}
		positions = nextPositions
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}
