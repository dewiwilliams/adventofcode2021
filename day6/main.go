package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := getData()
	fmt.Printf("Got data: %v\n", data)

	part1(data)
}
func part1(values []int) {
	workingSet := aggregateData(values)

	for i := 0; i < 80; i++ {
		workingSet = iterate(workingSet)
	}

	fmt.Printf("Total fish: %d\n", countFish(workingSet))
}
func countFish(values [9]int) int {
	result := 0

	for _, value := range values {
		result += value
	}

	return result
}
func iterate(values [9]int) [9]int {
	result := [9]int{}

	result[8] = values[0]
	result[7] = values[8]
	result[6] = values[7] + values[0]
	result[5] = values[6]
	result[4] = values[5]
	result[3] = values[4]
	result[2] = values[3]
	result[1] = values[2]
	result[0] = values[1]

	return result
}
func aggregateData(values []int) [9]int {
	result := [9]int{}

	for _, value := range values {
		result[value]++
	}

	return result
}
func getData() []int {
	result := []int{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	values := strings.Split(scanner.Text(), ",")

	for _, value := range values {
		value, err := strconv.Atoi(value)
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

	return result
}
