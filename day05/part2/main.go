package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"sort"
)

func main() {
	filePath := "../input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var originalRanges []Interval

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		bounds := strings.Split(line, "-")	
		low, _ := strconv.Atoi(bounds[0])
		high, _ := strconv.Atoi(bounds[1])
		originalRanges = append(originalRanges, Interval{Low: low, High: high})
	}

	sort.Slice(originalRanges, func(i, j int) bool {
		return originalRanges[i].Low < originalRanges[j].Low
	})	

	merged := []Interval{originalRanges[0]}

	for _, current := range originalRanges[1:] {
		last := &merged[len(merged) - 1]

		if last.High + 1 >= current.Low {
			last.High = max(last.High, current.High)
		} else {
			merged = append(merged, current)
		}
	}

	total := 0

	for _, current := range merged {
		total += current.High - current.Low + 1
	}

	fmt.Println(total)
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}

type Interval struct {
    Low, High int
}

