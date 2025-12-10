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

	scanner := bufio.NewScanner(file)
	tree := NewIntervalTree()

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		bounds := strings.Split(line, "-")	
		low, _ := strconv.Atoi(bounds[0])
		high, _ := strconv.Atoi(bounds[1])
		tree.Insert(low, high)
	}

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		ingredient, _ := strconv.Atoi(line)
		if tree.Contains(ingredient) {
			total++
		}
	}

	fmt.Println(total)
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}

type Interval struct {
    Low, High int
}

type Node struct {
    Interval
    MaxHigh int
    Left    *Node
    Right   *Node
}

type IntervalTree struct {
    root *Node
}

func NewIntervalTree() *IntervalTree {
    return &IntervalTree{}
}

func (it *IntervalTree) Insert(low, high int) {
    it.root = insert(it.root, low, high)
}

func insert(node *Node, low, high int) *Node {
    if node == nil {
        return &Node{
            Interval: Interval{Low: low, High: high},
            MaxHigh:  high,
        }
    }

    if low < node.Low {
        node.Left = insert(node.Left, low, high)
    } else {
        node.Right = insert(node.Right, low, high)
    }

    node.MaxHigh = node.High
    if node.Left != nil && node.Left.MaxHigh > node.MaxHigh {
        node.MaxHigh = node.Left.MaxHigh
    }
    if node.Right != nil && node.Right.MaxHigh > node.MaxHigh {
        node.MaxHigh = node.Right.MaxHigh
    }

    return node
}

func (it *IntervalTree) Contains(point int) bool {
    results := it.FindAll(point)
    return len(results) > 0
}

func (it *IntervalTree) FindAll(point int) []Interval {
    var results []Interval
    findAll(it.root, point, &results)
    return results
}

func findAll(node *Node, point int, results *[]Interval) {
	if node == nil {
			return
	}

	if point > node.MaxHigh {
			return
	}

	findAll(node.Left, point, results)

	if point >= node.Low && point <= node.High {
			*results = append(*results, node.Interval)
	}

	if node.Right != nil && point <= node.Right.MaxHigh {
			findAll(node.Right, point, results)
	}
}
