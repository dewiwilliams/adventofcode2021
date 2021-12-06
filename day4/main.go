package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type board struct {
	numbers [25]int
	state   [25]bool
}

func (b *board) numberCalled(number int) {
	for i := 0; i < 25; i++ {
		if b.numbers[i] == number {
			b.state[i] = true
		}
	}
}
func (b *board) isBingo() bool {
	return b.isHorizontalBingo() || b.isVerticalBingo()
}
func (b *board) isHorizontalBingo() bool {
	for y := 0; y < 5; y++ {
		if b.isBingoLine(y*5, 1) {
			return true
		}
	}

	return false
}
func (b *board) isVerticalBingo() bool {
	for x := 0; x < 5; x++ {
		if b.isBingoLine(x, 5) {
			return true
		}
	}

	return false
}
func (b *board) isBingoLine(start, stride int) bool {
	for i := 0; i < 5; i++ {
		if b.state[start+i*stride] == false {
			return false
		}
	}

	return true
}
func (b *board) getSumUnmarkedNumbers() int {
	result := 0

	for i := 0; i < 25; i++ {
		if !b.state[i] {
			result += b.numbers[i]
		}
	}

	return result
}

func main() {
	numbers, boards := getData()

	fmt.Printf("Got numbers: %v\n", numbers)

	part1(numbers, boards)
	part2(numbers, boards)
}

func remove(s []board, i int) []board {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func part2(numbers []int, boards []board) {

	workingSet := boards

	for _, number := range numbers {
		for true {
			if index := numberCalled(workingSet, number); index != -1 {

				if len(workingSet) == 1 {
					fmt.Printf("Part 2 answer: %d\n", boards[index].getSumUnmarkedNumbers()*number)
					return
				}

				workingSet = remove(workingSet, index)
			} else {
				break
			}
		}
	}

	fmt.Println("Impossible result")
	os.Exit(2)
}
func part1(numbers []int, boards []board) {

	for _, number := range numbers {
		if index := numberCalled(boards, number); index != -1 {
			fmt.Printf("Part 1 answer: %d\n", boards[index].getSumUnmarkedNumbers()*number)
			return
		}
	}

	fmt.Println("Impossible result")
	os.Exit(2)
}
func numberCalled(boards []board, number int) int {
	for i, _ := range boards {
		boards[i].numberCalled(number)
		if boards[i].isBingo() {
			return i
		}
	}

	return -1
}
func getData() ([]int, []board) {
	numbers := []int{}
	boards := []board{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	numbersPart := strings.Split(scanner.Text(), ",")
	for _, numberString := range numbersPart {
		value, err := strconv.Atoi(numberString)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		numbers = append(numbers, value)
	}

	//Empty line
	scanner.Scan()

	for true {
		newBoard := readBoard(scanner)
		if newBoard == nil {
			break
		}

		boards = append(boards, *newBoard)
	}

	return numbers, boards
}
func readBoard(scanner *bufio.Scanner) *board {
	b := board{}

	for i := 0; i < 5; i++ {
		if !scanner.Scan() {
			return nil
		}

		numbers := strings.Fields(scanner.Text())
		if len(numbers) != 5 {
			fmt.Printf("Wrong number of input numbers: %d\n", len(numbers))
			os.Exit(2)
		}
		for j := 0; j < 5; j++ {
			value, err := strconv.Atoi(numbers[j])
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}

			b.numbers[i*5+j] = value
		}
	}

	//Empty line
	scanner.Scan()

	return &b
}
