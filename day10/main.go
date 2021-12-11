package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := getData()

	part1(lines)
}
func part1(lines []string) {
	score := 0

	for _, line := range lines {
		score += getPart1Score(line)
	}

	fmt.Printf("Part 1 score: %d\n", score)
}
func getPart1Score(line string) int {
	stack := ""

	closingRunes := map[rune]rune{}
	closingRunes['('] = ')'
	closingRunes['['] = ']'
	closingRunes['{'] = '}'
	closingRunes['<'] = '>'

	scores := map[rune]int{}
	scores[')'] = 3
	scores[']'] = 57
	scores['}'] = 1197
	scores['>'] = 25137

	for _, char := range line {
		if char == '(' || char == '[' || char == '{' || char == '<' {
			stack += string(char)
			continue
		}

		if len(stack) == 0 {
			fmt.Printf("Empty stack: %s\n", line)
			os.Exit(2)
		}

		lastRune := []rune(stack[len(stack)-1:])
		if closingRunes[lastRune[0]] != char {
			return scores[char]
		}

		stack = stack[:len(stack)-1]
	}

	return 0
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

	return result
}
