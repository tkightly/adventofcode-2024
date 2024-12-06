package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"math"
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

	safeReportCount := 0

// break label to easily break out of an inner loop #NotAGoto
eachReport:

	for scanner.Scan() {
		report := strings.Split(scanner.Text(), " ")

		// we're going to loop through each interval, i.e. each adjacent pair of numbers (count of which is n-1, where n=count of all numbers)
		levelChanges := len(report) - 1

		// fmt.Printf("Debug: report: %s, levelChanges: %d\n", report, levelChanges)

		previousTrend := "unset"
		for i := 0; i < levelChanges; i++ {

			// work through each pair of numbers, called left and right
			trend := "unset"
			left, err := strconv.Atoi(report[i])
			if err != nil {
				// log.Printf("Error converting %s to integer: %v", report[i], err)
				continue eachReport
			}

			right, err := strconv.Atoi(report[i+1])
			if err != nil {
				// log.Printf("Error converting %s to integer: %v", report[i+1], err)
				continue eachReport
			}

			// fmt.Printf("  current change: %d, comparing %d and %d\n", i, left, right)

			if left > right {
				trend = "decreasing"
			}
			if right > left {
				trend = "increasing"
			}
			if left == right {
				trend = "equal"
			}

			interval := int(math.Abs(float64(left)-float64(right)))

			// fmt.Printf("    trend: %s\n", trend)

			// we don't know what the trend is until we look at the second pair of numbers and beyond, so i can't be 0
			if (i != 0 && trend != previousTrend) || trend == "equal" {
				// fmt.Printf("    breaking - trend changes\n")
				fmt.Printf("Report: %s: Invalid :(\n", report)
				continue eachReport
			} else if i != 0 && interval == 0 || interval > 3 {
				// fmt.Printf("    breaking - interval too great (%d, should be in 1,2,3)\n", interval)
				fmt.Printf("Report: %s: Invalid :(\n", report)
				continue eachReport

			}

			previousTrend = trend

		}

		fmt.Printf("Report: %s: Valid :)\n", report)
		safeReportCount++

	}
	fmt.Printf("Count of safe reports: %d\n", safeReportCount)
}
