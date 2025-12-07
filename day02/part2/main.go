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
	alreadySeen := make(map[int]struct{})
	// fmt.Printf("Testing range %s\n", rangeStr)

	for chunkSize := 1; chunkSize <= len(bounds[1])/2; chunkSize++ {
		chunkInt, _ := strconv.Atoi(bounds[0][:chunkSize])
		targetCheckSize := len(bounds[0])
		maxCheckSize := len(bounds[1])
		for targetCheckSize <= maxCheckSize {
			if targetCheckSize % chunkSize == 0 && targetCheckSize > 1 {
			  // fmt.Printf("testing targetCheckSize %d chunk Size %d, current %d\n", targetCheckSize, chunkSize, chunkInt)
				repeatedTimes := targetCheckSize / chunkSize
				for {
					repeatedValue := strconv.Itoa(chunkInt)

					if len(repeatedValue) > chunkSize {
						break
					}
					checkedValue, _ := strconv.Atoi(strings.Repeat(repeatedValue, repeatedTimes))
					if _, seen := alreadySeen[checkedValue]; seen {
						chunkInt++
						continue
					}
					alreadySeen[checkedValue] = struct{}{}
					if checkedValue > upperBound {
						break
					}
					if checkedValue >= lowerBound {
						// fmt.Println(checkedValue)
						totalInvalid += checkedValue
					} 
					chunkInt++
				}
			}
			chunkInt = int(math.Pow(10, float64(chunkSize - 1)))
			targetCheckSize++
		}
	}
	return totalInvalid
}

