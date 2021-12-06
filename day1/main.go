package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	data := getData()

	//Part 1
	result1 := 0
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			result1++
		}
	}

	fmt.Printf("Part 1 result: %d\n", result1)

	//Part 2
	result2 := 0
	for i := 3; i < len(data); i++ {
		value1 := data[i-3] + data[i-2] + data[i-1]
		value2 := data[i-2] + data[i-1] + data[i-0]

		if value1 < value2 {
			result2++
		}
	}
	fmt.Printf("Part 2 result: %d\n", result2)
}

func getData() []int {
	result := []int{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		value, err := strconv.Atoi(scanner.Text())
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
