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
