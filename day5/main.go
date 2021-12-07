package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type line struct {
	start [2]int
	end   [2]int
}

func (l *line) isHorizontal() bool {
	return l.start[1] == l.end[1]
}
func (l *line) isVertical() bool {
	return l.start[0] == l.end[0]
}
func (l *line) isDiagonal() bool {
	return l.start[0] != l.end[0] && l.start[1] != l.end[1]
}
func (l *line) applyHorizontalLineToBoard(board []int, stride int) {
	startx := l.start[0]
	endx := l.end[0]

	for x := startx; x <= endx; x++ {
		board[x+l.start[1]*stride]++
	}
}
func (l *line) applyVerticalLineToBoard(board []int, stride int) {
	starty := l.start[1]
	endy := l.end[1]

	for y := starty; y <= endy; y++ {
		board[l.start[0]+y*stride]++
	}
}
func (l *line) applyDiagonal(board []int, stride int) {
	step := 1
	if l.end[1] < l.start[1] {
		step = -1
	}

	startx := l.start[0]
	endx := l.end[0]
	limit := endx - startx
	starty := l.start[1]

	for i := 0; i <= limit; i++ {
		x := startx + i
		y := starty + i*step

		board[x+y*stride]++
	}
}
func (l *line) applyLineToBoard(board []int, stride int, applyDiagonals bool) {
	if l.isHorizontal() {
		l.applyHorizontalLineToBoard(board, stride)
	} else if l.isVertical() {
		l.applyVerticalLineToBoard(board, stride)
	} else if l.isDiagonal() {
		if applyDiagonals {
			l.applyDiagonal(board, stride)
		}
	} else {
		fmt.Println("Unrecognised line state")
		os.Exit(2)
	}
}

func main() {
	data := getData()

	part1(data)
	part2(data)
}

func part1(lines []line) {
	maximums := getBoardMaximums(lines)
	maximums[0]++
	maximums[1]++

	boardState := make([]int, maximums[0]*maximums[1])

	for _, line := range lines {
		line.applyLineToBoard(boardState, maximums[0], false)
	}

	fmt.Printf("Part 1 points: %d\n", countOverlapPoints(boardState))
}
func part2(lines []line) {
	maximums := getBoardMaximums(lines)
	maximums[0]++
	maximums[1]++

	boardState := make([]int, maximums[0]*maximums[1])

	for _, line := range lines {
		line.applyLineToBoard(boardState, maximums[0], true)
	}

	fmt.Printf("PArt 2 points: %d\n", countOverlapPoints(boardState))
}
func countOverlapPoints(board []int) int {
	result := 0

	for _, value := range board {
		if value > 1 {
			result++
		}
	}

	return result
}

func getBoardMaximums(lines []line) [2]int {
	x := 0
	y := 0

	for _, line := range lines {
		if line.start[0] > x {
			x = line.start[0]
		}
		if line.end[0] > x {
			x = line.end[0]
		}

		if line.start[1] > y {
			y = line.start[1]
		}
		if line.end[1] > y {
			y = line.end[1]
		}
	}

	return [2]int{x, y}
}

func getData() []line {
	result := []line{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		fields := strings.Fields(scanner.Text())

		line := line{}

		p1 := parseCoordinate(fields[0])
		p2 := parseCoordinate(fields[2])

		if p1[0] < p2[0] {
			line.start = p1
			line.end = p2
		} else if p1[0] > p2[0] {
			line.start = p2
			line.end = p1
		} else if p1[1] < p2[1] {
			line.start = p1
			line.end = p2
		} else {
			line.start = p2
			line.end = p1
		}

		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
func parseCoordinate(source string) [2]int {
	coordinates := strings.Split(source, ",")

	x, err := strconv.Atoi(coordinates[0])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	y, err := strconv.Atoi(coordinates[1])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	return [2]int{x, y}
}
