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

	fmt.Printf("Got values: %v", data)

	result := 0

	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			result++
		}
	}

	fmt.Printf("Result: %d\n", result)
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

		//fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
