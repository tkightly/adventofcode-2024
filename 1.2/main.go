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

	similarityScore := 0

	for index, leftNumber := range left {
		coincidenceCount := 0
		for _, rightNumber := range right {
			if leftNumber == rightNumber {
				coincidenceCount++
			}
		}
		fmt.Printf("debug: Number %d, value: %d, coincidenceCount: %d\n", index, leftNumber, coincidenceCount)

		similarityScore = similarityScore + (leftNumber * coincidenceCount)

	}

	fmt.Printf("Similarity Score: %d\n", similarityScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
