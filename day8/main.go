package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type dataLine struct {
	signalPatterns [10]string
	output         [4]string
}

func main() {
	data := getData()
	//fmt.Printf("Got data: %v\n", data)

	part1(data)
}

func part1(data []dataLine) {
	count := 0

	for _, item := range data {
		for i := 0; i < 4; i++ {
			length := len(item.output[i])
			if length == 2 || // Digit 1
				length == 4 || // Digit 4
				length == 3 || // Digit 7
				length == 7 { // Digit 8
				count++
			}
		}
	}

	fmt.Printf("Part 1 result: %d\n", count)
}

func getData() []dataLine {
	result := []dataLine{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		newLine := dataLine{
			signalPatterns: [10]string{
				fields[0],
				fields[1],
				fields[2],
				fields[3],
				fields[4],
				fields[5],
				fields[6],
				fields[7],
				fields[8],
				fields[9],
			},
			output: [4]string{
				fields[11],
				fields[12],
				fields[13],
				fields[14],
			},
		}

		result = append(result, newLine)
	}
	return result
}
