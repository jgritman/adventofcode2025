package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
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
	lineLength := len(first)
	numbersByPos := make([]int, lineLength)
	total := 0

	process := func(s string) {
		for i := 0; i < len(s); i++ {
			r := rune(s[i])
			if unicode.IsSpace(r) {
				continue
			}
			if r == '*' {
				subtotal := 1
				for i < len(numbersByPos) && numbersByPos[i] != 0 {
					subtotal *= numbersByPos[i]
					i++
				}
				total += subtotal
			} else if r == '+' {
				for i < len(numbersByPos) && numbersByPos[i] != 0 {
					total += numbersByPos[i]
					i++
				}
			} else {
				numbersByPos[i] = numbersByPos[i] * 10 + int(r - '0')
			}
		}
	}

	process(first)
	for scanner.Scan() {
		line := scanner.Text()
		process(line)
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}
