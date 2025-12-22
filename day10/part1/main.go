package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func createButton(length int, rawButtons string) uint16 {
	var mask uint16

	parts := strings.Split(rawButtons, ",")

	for _, p := range parts {
		pos, _ := strconv.Atoi(p)
		mask |= (1 << (length - 1 - pos))
	}
	return mask
}

func main() {
	filePath := "../input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	lightsReplacer := strings.NewReplacer(".", "0", "#", "1")
	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		machineLength := len(parts[0]) - 2
		rawLights := lightsReplacer.Replace(strings.Trim(parts[0], "[]"))
		lightsVal, _ := strconv.ParseUint(rawLights, 2, 16)
		lights := uint16(lightsVal)

		var buttons []uint16
		for i := 1; i < len(parts) -1; i++ {
			buttonTrimmed := strings.Trim(parts[i], "()")
			button := createButton(machineLength, buttonTrimmed)
			buttons = append(buttons, button)
		}

		// BFS for the solution
		presses := 1
		states := []uint16{0}
		visited := map[uint16]struct{} { 0: {} }
		LightCheckLoop:
		for {
			var newStates []uint16
			for _, state := range states {
				for _, button := range buttons {
					newState := state ^ button
					if newState == lights {
						total += presses
						break LightCheckLoop
					}
					if _, exists := visited[newState]; !exists {
						visited[newState] = struct{}{}
						newStates = append(newStates, newState)
					}
				}
			}
			states = newStates
			presses++
		}
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}
