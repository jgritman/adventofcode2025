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
	positions := map[int]int { startPos: 1 }

	for scanner.Scan() {
		line := scanner.Text()
		nextPositions := make(map[int]int)
		for pos := range positions {
			if line[pos] == '^' {
				leftSplit := pos - 1
				if leftSplit >= 0 {
					nextPositions[leftSplit] += positions[pos]
				} 
				rightSplit := pos + 1
				if rightSplit < len(line) {
					nextPositions[rightSplit] += positions[pos]
				}
			} else {
				nextPositions[pos] += positions[pos]
			} 
		}
		positions = nextPositions
	}

	total := 0
	for pos := range positions {
		total += positions[pos]
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}
