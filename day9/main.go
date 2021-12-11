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

	fmt.Printf("Got data: %v\n", grid)
	fmt.Printf("Got width: %d\n", width)

	part1(grid, width)
}
func part1(grid []int, width int) {
	fmt.Printf("Part 1 risk level: %d\n", getRiskLevel(grid, width))
}
func getRiskLevel(grid []int, width int) int {
	height := len(grid) / width
	risklevel := 0

	fmt.Printf("Got height: %d\n", height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			index := x + y*width

			if x > 0 && grid[index-1] <= grid[index] {
				continue
			}
			if x < width-1 && grid[index+1] <= grid[index] {
				continue
			}
			if y > 0 && grid[index-width] <= grid[index] {
				continue
			}
			if y < height-1 && grid[index+width] <= grid[index] {
				continue
			}

			risklevel += grid[index] + 1
		}
	}

	return risklevel
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
