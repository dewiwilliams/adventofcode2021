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

	fmt.Printf("Got data: %s\n", data)

	position := 0
	depth := 0

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

		if pieces[0] == "forward" {
			position += v
		} else if pieces[0] == "up" {
			depth -= v
		} else if pieces[0] == "down" {
			depth += v
		}
	}

	fmt.Printf("Part 1 result: %d\n", position*depth)
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
