package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "regexp"
	// "strconv"
 // "test"
)

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

	fmt.Println(lines)

	// create the grid
	// define directions

	// the xmas word

	// loop each row

	// loop column

	// loop direction

	// loop characters within the word

	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	directions := [][2]int{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}

	word := "XMAS"
	output := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			for _, dir := range directions {
				// fmt.Printf("Row: %d, col: %d, direction: %s\n", row, col, dir)

				rowIndex := dir[0]
				colIndex := dir[1]
				isWord := true

				for charIndex := 0; charIndex < len(word); charIndex++ {

					rowOffset := row + (rowIndex * charIndex)
					colOffset := col + (colIndex * charIndex)

					if rowOffset < 0 || rowOffset >= len(grid) || colOffset < 0 || colOffset >= len(grid[row]) {
						isWord = false
						break
					}

					if grid[rowOffset][colOffset] != rune(word[charIndex]) {
						isWord = false
						break
					}
				}

				if isWord == true {
					output++
				}

			}

		}
	}

	fmt.Println("Output Day 4 Part 1", output)

}
