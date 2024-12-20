package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	var rules [][2]int
	var updates [][]int

	for scanner.Scan() {

		matched, _ := regexp.MatchString(`^\d*\|\d*$`, scanner.Text())

		if matched {

			var intValues [2]int
			strValues := strings.Split(scanner.Text(), "|")

			for i, value := range strValues {

				intValues[i], _ = strconv.Atoi(strings.TrimSpace(value))

			}

			rules = append(rules, intValues)
			continue
		}

		matched, _ = regexp.MatchString(`,`, scanner.Text())

		if matched {

			var intValues []int
			strValues := strings.Split(scanner.Text(), ",")

			for _, value := range strValues {

				value, _ := strconv.Atoi(strings.TrimSpace(value))
				intValues = append(intValues, value)
			}

			updates = append(updates, intValues)
		}

	}

	// fmt.Printf("Rules: %v\n", rules)
	// fmt.Printf("Updates: %v\n", updates)

	output := 0
	ruleValid := true

	for _, update := range updates {
		// fmt.Printf("This update: %v\n", update)
		for _, rule := range rules {
			ruleValid = true
			// fmt.Printf("  This rule: %v\n", rule)
			leftIndex := findIndex(update, rule[0])
			if leftIndex == -1 {
				// fmt.Printf("  This rule is disregarded as %v doesn't appear in the update %v\n", rule[0], update)
				// fmt.Printf("    leftIndex: %d\n", leftIndex)
				continue
			}

			rightIndex := findIndex(update, rule[1])

			if rightIndex == -1 {
				// fmt.Printf("  This rule is disregarded as %v doesn't appear in the update %v\n", rule[1], update)
				// fmt.Printf("    rightIndex: %d\n", rightIndex)
				continue
			}

			if leftIndex >= rightIndex {
				// fmt.Printf("  This rule invalid because rightIndex (%v) appears equal to or before leftIndex (%v)\n", rightIndex, leftIndex)
				ruleValid = false
				break
			}

			// fmt.Printf("  This rule valid :) \n")

			// determine if rule is valid
			// if rule invalid, break and set var saying the update is invalud

		}
		if ruleValid == true {
			// fmt.Printf(" This update is in the right order :)\n")

			len := len(update)
			middleIndex := math.Floor(float64(len) / 2)
			// fmt.Println(middleIndex)
			output = output + update[int(middleIndex)]
		} else {
			// fmt.Printf(" This update is not in the right order :(\n")
		}

		// if the update is invalid, break
		// otherwise add the middle number to the output array
	}
	fmt.Printf("Output: %v", output)
}

func findIndex(arr []int, target int) int {
	for i, value := range arr {
		if value == target {
			return i
		}
	}
	return -1
}
