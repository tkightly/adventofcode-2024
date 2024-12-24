package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func moveGuard(dir string, location [2]int, grid [][]rune) (string, [2]int, [][]rune, bool, bool) {
	newLocation := location

	switch dir {
	case "up":
		newLocation[0] = location[0] - 1
	case "right":
		newLocation[1] = location[1] + 1
	case "down":
		newLocation[0] = location[0] + 1
	case "left":
		newLocation[1] = location[1] - 1
	}


	if newLocation[0] < 0 || newLocation[0] >= len(grid) || newLocation[1] < 0 || newLocation[1] >= len(grid[0]) {
		return dir, newLocation, grid, true, false // true == out of bounds
	}

	if grid[newLocation[0]][newLocation[1]] == '#' {
		return dir, location, grid, false, true // true == obstacle
	}

	grid[newLocation[0]][newLocation[1]] = 'X'
	return dir, newLocation, grid, false, false // true == in bounds

}
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}


	for _, line := range grid {
		fmt.Printf("%s\n", string(line))
	}

	var guardLocation [2]int
	var guardDir string

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '^' {
				guardLocation = [2]int{row, col}
				guardDir = "up"
				fmt.Printf("Found guard at (row %d, col %d), going %s\n", guardLocation[0], guardLocation[1], guardDir)
			}
		}
	}

	var obstacle bool = false
	var outOfBounds bool = false

	for i := 0; obstacle == false && outOfBounds == false; i++ {

		guardDir, guardLocation, grid, outOfBounds, obstacle = moveGuard(guardDir, guardLocation, grid)

		if obstacle == true {
			obstacle = false
			switch guardDir {
			case "up":
				guardDir = "right"
			case "right":
				guardDir = "down"
			case "down":
				guardDir = "left"
			case "left":
				guardDir = "up"
			}

			grid[guardLocation[0]][guardLocation[1]] = '@'

			continue

		}

		if outOfBounds == true {

			break
		}

	}

	locations := 0
	for _, line := range grid {
		fmt.Printf("%s\n", string(line))
		for _, char := range line {
			if char == 'X' || char == '@' || char == '^' {
				locations++
			}
		}

	}

	fmt.Println("Output Day 6 Part 1: ", locations)

}
