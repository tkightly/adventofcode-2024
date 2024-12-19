package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// close the input file at the end

	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0

	for scanner.Scan() {

		r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
		instructions := r.FindAllString(scanner.Text(), -1)

		for _, instruction := range instructions {
			r, _ := regexp.Compile(`\d{1,3}`)
			values := r.FindAllString(instruction, -1)

			leftValue, err := strconv.Atoi(values[0])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			rightValue, err := strconv.Atoi(values[1])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			product := leftValue * rightValue

			total = total + product

		}
		fmt.Printf("Running total: %d\n", total)
	}

	fmt.Printf("Final Total: %d\n", total)
}
