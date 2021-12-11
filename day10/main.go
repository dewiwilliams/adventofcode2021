package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	lines := getData()

	part1(lines)
	part2(lines)
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
func part2(lines []string) {
	scores := []int{}

	for _, line := range lines {
		score := getPart1Score(line)
		if score != 0 {
			continue
		}

		scores = append(scores, getPart2Score(line))
	}

	sort.Ints(scores)
	index := (len(scores) - 1) / 2

	fmt.Printf("Part 2 score: %d\n", scores[index])
}
func getPart2Score(line string) int {
	stack := ""

	closingRunes := map[rune]rune{}
	closingRunes['('] = ')'
	closingRunes['['] = ']'
	closingRunes['{'] = '}'
	closingRunes['<'] = '>'

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
			fmt.Printf("Corrupted data line: %s\n", line)
			os.Exit(2)
		}

		stack = stack[:len(stack)-1]
	}

	return calculateIncompleteLineScore(stack)
}
func calculateIncompleteLineScore(s string) int {
	result := 0

	scores := map[rune]int{}
	scores['('] = 1
	scores['['] = 2
	scores['{'] = 3
	scores['<'] = 4

	for i := len(s) - 1; i >= 0; i-- {
		result *= 5
		result += scores[rune(s[i])]
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

	return result
}
