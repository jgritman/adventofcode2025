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
	X, Y int
}

type Rect struct {
	Min, Max Point
}

func (r Rect) Overlaps(other Rect) bool {
	return r.Min.X < other.Max.X &&
		r.Max.X > other.Min.X &&
		r.Min.Y < other.Max.Y &&
		r.Max.Y > other.Min.Y
}

func NewRectFromPoints(p1, p2 Point) Rect {
	return Rect{
		Min: Point{
			X: min(p1.X, p2.X),
			Y: min(p1.Y, p2.Y),
		},
		Max: Point{
			X: max(p1.X, p2.X),
			Y: max(p1.Y, p2.Y),
		},
	}
}

func (r Rect) Area() int {
  return (r.Max.X - r.Min.X + 1) * (r.Max.Y - r.Min.Y + 1)
}

// wire "fence" is a 1px rectange outside the border
func wireRects(tiles []Point) []Rect {
	isCW := isClockwise(tiles)
	var wires []Rect

	addWire := func(p1, p2 Point) {
		var wireP1, wireP2 Point
		if p1.X == p2.X {
			minY, maxY := min(p1.Y, p2.Y) + 1, max(p1.Y, p2.Y) - 1
			up := p2.Y < p1.Y
			wireX := p1.X + 1
			if up == isCW {
				wireX = p1.X - 1
			}
			wireP1 = Point{Y: minY, X: wireX }
			wireP2 = Point{Y: maxY, X: wireX }
		} else {
			minX, maxX := min(p1.X, p2.X) + 1, max(p1.X, p2.X) - 1
			right := p2.X > p1.X
			wireY := p1.Y + 1
			if right == isCW {
				wireY = p1.Y - 1
			}
			wireP1 = Point{Y: wireY, X: minX }
			wireP2 = Point{Y: wireY, X: maxX }
		}
		wire := NewRectFromPoints(wireP1, wireP2)
	// fmt.Printf("SEGMENT: %+v -> %+v | WIRE: %+v\n", p1, p2, wire)
		wires = append(wires, wire)
	}

	n := len(tiles)
	for i := 0; i < n - 1; i++ {
		addWire(tiles[i], tiles[i + 1])
	}
	addWire(tiles[n - 1], tiles[0])
	return wires
}

// "shoelace" fomula
func isClockwise(tiles []Point) bool {
	sum := 0
	n := len(tiles)
	for i := 0; i < n - 1; i++ {
    p1, p2 := tiles[i], tiles[i + 1]
    sum += (p2.X - p1.X) * (p2.Y + p1.Y)
	}
	pLast, pFirst := tiles[n - 1], tiles[0]
	sum += (pFirst.X - pLast.X) * (pFirst.Y + pLast.Y)
	return sum < 0
}

func hitsWire(r Rect, wires []Rect) bool {
	for _, wire := range wires {
		if r.Overlaps(wire) {
      // fmt.Printf("REJECTED: Candidate %+v hit Wire %+v\n", r, wire)
			return true
		}
	}
	return false
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
		x, _ := strconv.Atoi(tokens[0])
		y, _ := strconv.Atoi(tokens[1])
		redTiles = append(redTiles, Point{X: x, Y: y})
	}

	maxArea := 0
	wires := wireRects(redTiles)

	for i, tile1 := range redTiles {
		for j := i + 1; j < len(redTiles); j++ {
			toCheck := NewRectFromPoints(tile1, redTiles[j])
			area := toCheck.Area()
			// fmt.Println(tile1)
			// fmt.Println(redTiles[j])
			// fmt.Println(area)
			if area > maxArea {
				if !hitsWire(toCheck, wires) {
				 maxArea = area
				}
			}
		}
	}

	fmt.Println(maxArea)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}
