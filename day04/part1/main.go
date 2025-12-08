package main

import (
	"fmt"
	"aoc/common"
)

const RollChar = '@' 

func main() {
	filePath := "../input.txt"

	layout, _ := common.ReadFileToChar2D(filePath) 
	
	total := 0

	for y := 0; y < len(layout); y++ {
		for x := 0; x < len(layout[y]); x++ {
			currentChar, _ := common.Get(layout, y, x)
			if currentChar == RollChar {
			  // fmt.Printf("checking %d %d\n", y, x)	

				neighborChars := common.NeighborChars(layout, y, x)
				// fmt.Println(neighborChars)
				count := 0
				for _, neighbor := range neighborChars {
					if neighbor == RollChar {
						count++
					}
				}
				// fmt.Println(count)
				if count < 4 {
					// fmt.Println("adding")
					total++
				}
			}
		}
	}
	fmt.Println(total)
}
