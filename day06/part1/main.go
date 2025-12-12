package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	filePath := "../input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	var addArray []int
	var multArray []int
	total := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		for i, token := range tokens {
			if token == "*" {
				// fmt.Println(multArray[i])
				total += multArray[i]
			} else if token == "+" {
				// fmt.Println(addArray[i])
				total += addArray[i]
			} else {
				tokenInt, _ := strconv.Atoi(token)
				if len(addArray) <= i {
					addArray = append(addArray, tokenInt)
					multArray = append(multArray, tokenInt)
				} else {
					addArray[i] += tokenInt
					multArray[i] *= tokenInt
				}
			}
		}
		// fmt.Println(addArray)
		// fmt.Println(multArray)
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}
