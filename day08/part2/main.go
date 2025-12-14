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

type Junction struct {
	X, Y, Z int
	CircuitPos int
}

type JunctionDistance struct {
	A, B *Junction
	DistanceSq int
}

func dumpJunction(j *Junction) {
    fmt.Printf("%+v\n", j) 
}

func EuclideanDistance(a, b *Junction) JunctionDistance {
	dx :=a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	distanceSq := dx*dx + dy*dy + dz*dz
	return JunctionDistance { A: a, B: b, DistanceSq: distanceSq }
}

func main() {
	filePath := "../input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var junctions []*Junction

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		x, _ := strconv.Atoi(tokens[0])
		y, _ := strconv.Atoi(tokens[1])
		z, _ := strconv.Atoi(tokens[2])
		junction := &Junction{X: x, Y: y, Z: z, CircuitPos: -1}
		junctions = append(junctions, junction)
	}

	var distances []JunctionDistance
	for i, junction := range junctions {
		for j := i + 1; j < len(junctions); j++ {
			distance := EuclideanDistance(junction, junctions[j]) 
			distances = append(distances, distance)
		}
	}

	sort.Slice(distances, func(i, j int) bool { return distances[i].DistanceSq < distances[j].DistanceSq })

	circuits := make([][]*Junction, 0)
	currentDistancePos := -1

	for len(circuits) == 0 || len(circuits[0]) != len(junctions) {
		currentDistancePos++
		distance := distances[currentDistancePos]
		// dumpJunction(distance.A)
		// dumpJunction(distance.B)
		if distance.A.CircuitPos == -1 && distance.B.CircuitPos == -1 {
		  // create a new ciruit if neither connected
			// fmt.Println("creating new circuit")
			circuitJunctions := []*Junction { distance.A, distance.B }
			circuits = append(circuits, circuitJunctions)
			newCircuitPos := len(circuits) - 1
			distance.A.CircuitPos = newCircuitPos
			distance.B.CircuitPos = newCircuitPos
		} else if distance.A.CircuitPos == -1 {
			circuitPos := distance.B.CircuitPos
			circuits[circuitPos] = append(circuits[circuitPos], distance.A)
			distance.A.CircuitPos = circuitPos
		} else if distance.B.CircuitPos == -1 {
			circuitPos := distance.A.CircuitPos
			circuits[circuitPos] = append(circuits[circuitPos], distance.B)
			distance.B.CircuitPos = circuitPos
		} else if distance.A.CircuitPos != distance.B.CircuitPos {
			// always copy to the lowest bound circuit
			distanceAPos := min(distance.A.CircuitPos, distance.B.CircuitPos)
			distanceBPos := max(distance.A.CircuitPos, distance.B.CircuitPos)
			for _, junction := range circuits[distanceBPos] {
				circuits[distanceAPos] = append(circuits[distanceAPos], junction)
				junction.CircuitPos = distanceAPos
			} 
			circuits[distanceBPos] = []*Junction {}
		}
		// fmt.Println(circuits)
	}

	lastDistance := distances[currentDistancePos]
	result := lastDistance.A.X * lastDistance.B.X
	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}
