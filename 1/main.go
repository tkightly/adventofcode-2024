package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// close the input file at the end
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// left side set of numbers
	left := []int{}

	// right side set of numbers
	right := []int{}

	for scanner.Scan() {
		// left and right lists are separated by three spaces :shrug:
		line := strings.Split(scanner.Text(), "   ")

		// strconv.Atoi() converts from ascii to integer, hence the name
		leftValue, err := strconv.Atoi(line[0])
		if err != nil {
			log.Printf("Error converting %s to integer: %v", line[0], err)
			continue
		}

		rightValue, err := strconv.Atoi(line[1])
		if err != nil {
			log.Printf("Error converting %s to integer: %v", line[1], err)
			continue
		}

		left = append(left, leftValue)
		right = append(right, rightValue)
	}

	// sorts ascending
	sort.Ints(left)
	sort.Ints(right)

	runningDiff := 0

	// feels dirty but each list should have the same number of numbers... emphasis on should
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]

		// we're working out the difference between the two numbers, so if it's negative (i.e. when the larger number is subtracted by the smaller number) we need to invert the sign
		if diff < 0 {
			diff = diff * -1
		}

		// left this in for debug
		// fmt.Printf("Line: %d Diff: %d, left: %d, right: %d\n", i+1, diff, left[i], right[i])

		runningDiff = runningDiff + diff
	}

	fmt.Printf("Total difference is: %d\n", runningDiff)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
