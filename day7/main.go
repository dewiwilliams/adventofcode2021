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
	//fmt.Printf("Got data: %v\n", data)

	part1(data)
	//part2(data)
}

func part1(values []int) {
	minimum := calculateFuel(values, 0)

	limit := max(values)

	for i := 1; i < limit; i++ {
		cost := calculateFuel(values, i)
		if cost < minimum {
			minimum = cost
		}
	}

	fmt.Printf("Part 1 fuel cost: %d\n", minimum)
}

func calculateFuel(values []int, position int) int {
	result := 0

	for _, value := range values {
		result += abs(value - position)
	}

	return result
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func max(values []int) int {
	result := 0

	for i, e := range values {
		if i == 0 || e > result {
			result = e
		}
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
