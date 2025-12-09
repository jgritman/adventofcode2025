package common

import (
    "bufio"
    "os"
)

type Coord struct {
	Row, Col int
}

func NewCoord(row, col int) Coord {
    return Coord{Row: row, Col: col}
}

var Directions = [8][2]int{
    {-1, -1},
    {-1,  0},
    {-1,  1},
    { 0, -1},
    { 0,  1},
    { 1, -1},
    { 1,  0},
    { 1,  1},
}

func ReadFileToChar2D(filepath string) ([][]rune, error) {
    file, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var result [][]rune
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        line := scanner.Text()
        row := make([]rune, len(line))
        for i, char := range line {
            row[i] = rune(char)
        }
        result = append(result, row)
    }
    
    return result, scanner.Err()
}

func Get(grid [][]rune, row, col int) (rune, bool) {
    if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
        return 0, false
    }
    return grid[row][col], true
}

func NeighborChars(grid [][]rune, row, col int) []rune {
    var chars []rune
    
    for i := 0; i < 8; i++ {
        dr, dc := Directions[i][0], Directions[i][1]
        newRow, newCol := row+dr, col+dc
        
        if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[newRow]) {
            chars = append(chars, grid[newRow][newCol])
        }
    }
    
    return chars
}
