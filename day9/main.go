package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	grid, width := getData()

	part1(grid, width)
	part2(grid, width)
}
func part1(grid []int, width int) {
	fmt.Printf("Part 1 risk level: %d\n", getRiskLevel(grid, width))
}
func part2(grid []int, width int) {
	basinSizes := getBasinSizes(grid, width)
	sort.Ints(basinSizes)

	largestBasins := basinSizes[len(basinSizes)-3:]
	result := largestBasins[0] * largestBasins[1] * largestBasins[2]

	fmt.Printf("Part 2 result: %d\n", result)
}
func getRiskLevel(grid []int, width int) int {
	risklevel := 0

	lowPoints := getLowPoints(grid, width)
	for _, point := range lowPoints {
		risklevel += grid[point] + 1
	}

	return risklevel
}
func getLowPoints(grid []int, width int) []int {
	result := []int{}

	height := len(grid) / width

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

			result = append(result, index)
		}
	}

	return result
}
func getBasinSizes(grid []int, width int) []int {
	result := []int{}

	lowPoints := getLowPoints(grid, width)
	for _, point := range lowPoints {
		result = append(result, getBasinSize(grid, width, point))
	}

	return result
}
func getBasinSize(grid []int, width, lowpoint int) int {
	pointsToProcess := map[int]bool{}
	pointsToProcess[lowpoint] = true

	height := len(grid) / width

	pointsInBasin := map[int]bool{}
	pointsInBasin[lowpoint] = true

	for len(pointsToProcess) > 0 {
		key := getFirstKey(pointsToProcess)
		delete(pointsToProcess, key)

		neighbours := getNeighbours(key, width, height)
		for _, neighbour := range neighbours {
			if grid[neighbour] == 9 {
				continue
			}
			if grid[neighbour] > grid[key] {
				pointsInBasin[neighbour] = true
				pointsToProcess[neighbour] = true
			}
		}

	}

	return len(pointsInBasin)
}
func getFirstKey(m map[int]bool) int {
	for k := range m {
		return k
	}
	return -1
}
func getNeighbours(point, width, height int) []int {
	result := []int{}

	x := point % width
	y := point / width

	if x > 0 {
		result = append(result, point-1)
	}
	if x < width-1 {
		result = append(result, point+1)
	}
	if y > 0 {
		result = append(result, point-width)
	}
	if y < height-1 {
		result = append(result, point+width)
	}

	return result
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
