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
    fmt.Printf("%+v\n", j) // prints field names and values
}

func EuclideanDistance(a, b *Junction) JunctionDistance {
	dx :=a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	distanceSq := dx*dx + dy*dy + dz*dz
	return JunctionDistance { A: a, B: b, DistanceSq: distanceSq }
}

func main() {
	noOfClosest := 1000
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

	for i := 0; i < noOfClosest; i++ {
		distance := distances[i]
		// dumpJunction(distance.A)
		// dumpJunction(distance.B)
		if distance.A.CircuitPos == -1 && distance.B.CircuitPos == -1 {
		  // create a new ciruit if neither connected
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
			// move all circuit B junctions to Junction A
			distanceAPos := distance.A.CircuitPos
			distanceBPos := distance.B.CircuitPos
			for _, junction := range circuits[distanceBPos] {
				circuits[distanceAPos] = append(circuits[distanceAPos], junction)
				junction.CircuitPos = distanceAPos
			} 
			circuits[distanceBPos] = []*Junction {}
		}
		// fmt.Println(circuits)
	}

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})

	total := 1
	for i := 0; i < 3; i++ {
		total *= len(circuits[i])
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}
