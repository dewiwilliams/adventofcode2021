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
	grid, width := getData()

	part1(grid, width)
}
func part1(grid []int, width int) {
	flashes := 0
	currentGrid := grid

	for i := 0; i < 100; i++ {
		newGrid, flashCount := iterate(currentGrid, width)

		currentGrid = newGrid
		flashes += flashCount
	}

	fmt.Printf("Part 1 number of flashes: %d", flashes)
}

func iterate(grid []int, width int) ([]int, int) {
	flashes := 0
	flashed := map[int]bool{}

	pointsToProcess := map[int]int{}
	for index, _ := range grid {
		pointsToProcess[index] = 1
	}

	for len(pointsToProcess) > 0 {
		key := getFirstKey(pointsToProcess)
		delete(pointsToProcess, key)
	}

	return grid, flashes
}
func getNighbours(cell, width int) []int {
	result := []int{}

	return result
}
func getFirstKey(m map[int]int) int {
	for k := range m {
		return k
	}
	return -1
}
func getData() ([]int, int) {
	result := []int{}
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

		numbers := strings.Split(text, "")
		for i := 0; i < len(numbers); i++ {
			value, err := strconv.Atoi(numbers[i])
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}

			result = append(result, value)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result, maxLength
}
