package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

type Point struct {
	Y, X int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Area(p1, p2 Point) int {
	height := Abs(p1.Y - p2.Y)  + 1
	width := Abs(p1.X - p2.X) + 1
	return height * width
}

func main() {
	filePath := "../input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var redTiles []Point

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		y, _ := strconv.Atoi(tokens[0])
		x, _ := strconv.Atoi(tokens[1])
		redTiles = append(redTiles, Point{X: x, Y: y})
	}

	maxArea := 0

	for i, tile1 := range redTiles {
		for j := i + 1; j < len(redTiles); j++ {
			area := Area(tile1, redTiles[j])
			// fmt.Println(tile1)
			// fmt.Println(redTiles[j])
			// fmt.Println(area)
			maxArea = max(maxArea, area)
		}
	}

	fmt.Println(maxArea)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}
