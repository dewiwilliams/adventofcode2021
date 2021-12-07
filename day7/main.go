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

	part1(data)
	part2(data)
}

func part2(values []int) {
	minimum := calculateGeometericFuelTotal(values, 0)

	limit := max(values)

	for i := 1; i < limit; i++ {
		cost := calculateGeometericFuelTotal(values, i)
		if cost < minimum {
			minimum = cost
		}
	}

	fmt.Printf("Part 2 fuel cost: %d\n", minimum)
}
func part1(values []int) {
	// Instead of linearly scanning the full search space, start at the median
	// and search either side from there until the cost is increasing.
	minimum := calculateLinearFuelTotal(values, 0)

	limit := max(values)

	for i := 1; i < limit; i++ {
		cost := calculateLinearFuelTotal(values, i)
		if cost < minimum {
			minimum = cost
		}
	}

	fmt.Printf("Part 1 fuel cost: %d\n", minimum)
}
func calculateGeometericFuelTotal(values []int, position int) int {
	result := 0

	for _, value := range values {
		result += calculateGeometricFuelCost(abs(value - position))
	}

	return result
}
func calculateGeometricFuelCost(distance int) int {
	//Sum of arithmetic sequence
	return ((distance + 1) * distance) / 2
}
func calculateLinearFuelTotal(values []int, position int) int {
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
