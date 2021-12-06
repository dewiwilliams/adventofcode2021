package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type dataItem struct {
	command string
	value   int
}

func main() {
	data := getData()
	parsedData := parseData(data)

	part1(parsedData)
	part2(parsedData)
}
func part2(data []dataItem) {
	position := 0
	aim := 0
	depth := 0

	for _, item := range data {
		if item.command == "forward" {
			position += item.value
			depth += aim * item.value
		} else if item.command == "up" {
			aim -= item.value
		} else if item.command == "down" {
			aim += item.value
		}
	}

	fmt.Printf("Part 1 result: %d\n", position*depth)
}
func part1(data []dataItem) {
	position := 0
	depth := 0

	for _, item := range data {
		if item.command == "forward" {
			position += item.value
		} else if item.command == "up" {
			depth -= item.value
		} else if item.command == "down" {
			depth += item.value
		}
	}

	fmt.Printf("Part 1 result: %d\n", position*depth)
}
func parseData(data []string) []dataItem {
	result := []dataItem{}

	for _, value := range data {
		pieces := strings.Fields(value)
		if len(pieces) != 2 {
			fmt.Printf("Unexpected data: %s\n", value)
			os.Exit(2)
		}

		v, err := strconv.Atoi(pieces[1])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		item := dataItem{
			command: pieces[0],
			value:   v,
		}
		result = append(result, item)
	}

	return result
}
func getData() []string {
	result := []string{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
