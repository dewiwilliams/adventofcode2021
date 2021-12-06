package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	data, maxLength := getData()

	fmt.Printf("Max length: %d\n", maxLength)

	part1(data, maxLength)
	part2(data, maxLength)
}
func part2(data []int64, maxLength int) {

	oxygenRating := getOxygenRating(data, maxLength)
	scrubberRating := getScrubberRating(data, maxLength)

	fmt.Printf("Part 2 result: %d\n", oxygenRating*scrubberRating)
}
func getScrubberRating(data []int64, maxLength int) int64 {
	workingSet := data
	for i := 0; i < maxLength; i++ {
		bit := maxLength - i - 1
		score := getDataSetScore(workingSet, bit)
		workingSet = filterDataSet(workingSet, bit, score < 0)

		if len(workingSet) == 1 {
			return workingSet[0]
		}
	}

	fmt.Println("Scrubber: Impossible result")
	os.Exit(2)
	return int64(0)
}
func getOxygenRating(data []int64, maxLength int) int64 {
	workingSet := data
	for i := 0; i < maxLength; i++ {
		bit := maxLength - i - 1
		score := getDataSetScore(workingSet, bit)
		workingSet = filterDataSet(workingSet, bit, score >= 0)

		if len(workingSet) == 1 {
			return workingSet[0]
		}
	}

	fmt.Println("Oxygen: Impossible result")
	os.Exit(2)
	return int64(0)
}
func filterDataSet(data []int64, bit int, state bool) []int64 {

	result := []int64{}

	testValue := int64(1 << bit)

	for _, dataPoint := range data {
		hasBit := (dataPoint & testValue) > 0
		if hasBit == state {
			result = append(result, dataPoint)
		}
	}

	return result
}
func getDataSetScore(data []int64, bit int) int {
	score := 0

	testValue := int64(1 << bit)
	for _, dataPoint := range data {
		if (dataPoint & testValue) > 0 {
			score++
		} else {
			score--
		}
	}

	return score
}
func part1(data []int64, maxLength int) {

	scores := make([]int, maxLength)

	for _, dataPoint := range data {
		for i := 0; i < maxLength; i++ {
			if dataPoint&1 > 0 {
				scores[i]++
			} else {
				scores[i]--
			}

			dataPoint >>= 1
		}
	}

	gamma := 0
	epsilon := 0

	for i := 0; i < maxLength; i++ {
		if scores[i] > 0 {
			gamma |= 1 << i
		} else {
			epsilon |= 1 << i
		}
	}

	fmt.Printf("Result: %d\n", gamma*epsilon)
}
func getData() ([]int64, int) {
	result := []int64{}
	maxLength := 0

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		text := scanner.Text()
		if len(text) > maxLength {
			maxLength = len(text)
		}

		value, err := strconv.ParseInt(text, 2, 64)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		result = append(result, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result, maxLength
}
