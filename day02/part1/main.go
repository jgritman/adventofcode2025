package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"math"
)

func main() {
	filePath := "../input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		line := scanner.Text()
		invalidTotal := 0
		ranges := strings.Split(line, ",")	
		for _, rangeStr := range ranges {
			invalidTotal += handleRange(rangeStr)
		}
		fmt.Println(invalidTotal)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}

func handleRange(rangeStr string) (int) {
	bounds := strings.Split(rangeStr, "-")	
	lowerBound, _ := strconv.Atoi(bounds[0])
	upperBound, _ := strconv.Atoi(bounds[1])
  totalInvalid := 0 
	
	currentHalf := 0
	lowerLength := len(bounds[0])
	if lowerLength % 2 == 0 {
		currentHalf, _ = strconv.Atoi(firstHalf(bounds[0]))
	} else {
		exponent := lowerLength / 2
		currentHalf = int(math.Pow(10, float64(exponent)))
	}

	for {
		currentHalfStr := strconv.Itoa(currentHalf)
		testValue, _ := strconv.Atoi(currentHalfStr + currentHalfStr)
		//fmt.Println(testValue)
		if testValue > upperBound {
			break
		}
		currentHalf++
		if testValue >= lowerBound {
			//fmt.Println(testValue)
			totalInvalid += testValue
		}
	}
	return totalInvalid
}

func firstHalf(s string) string {
    length := len(s)
    half := length / 2
    return s[:half]
}
