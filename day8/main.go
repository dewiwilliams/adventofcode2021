package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type dataLine struct {
	signalPatterns [10]string
	output         [4]string
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func runtimeAssert(state bool) {
	if state {
		return
	}

	fmt.Println("Assertion failed")
	os.Exit(2)
}

func main() {
	data := getData()

	part1(data)
	part2(data)
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
func part2(data []dataLine) {
	result := 0

	for _, item := range data {
		result += decodeLine(item)
	}

	fmt.Printf("Part 2 result: %d\n", result)
}
func decodeLine(data dataLine) int {

	signalMapping, _ := getMappings(data)

	mapping := invertSignalMapping(signalMapping)

	return mapping[data.output[0]]*1000 +
		mapping[data.output[1]]*100 +
		mapping[data.output[2]]*10 +
		mapping[data.output[3]]*1
}
func invertSignalMapping(mapping map[int]string) map[string]int {
	result := map[string]int{}

	for index, value := range mapping {
		result[value] = index
	}

	return result
}
func getMappings(data dataLine) (map[int]string, map[int]string) {
	/*
		 Segment positions

		 0000
		1    2
		1    2
		 3333
		4    5
		4    5
		 6666
	*/

	segmentMapping := map[int]string{}
	signalMapping := map[int]string{}

	findSignalOne(data, &signalMapping, &segmentMapping)
	findSignalFour(data, &signalMapping, &segmentMapping)
	findSignalSeven(data, &signalMapping, &segmentMapping)
	findSignalEight(data, &signalMapping, &segmentMapping)
	findSignalNine(data, &signalMapping, &segmentMapping)
	findSignalThree(data, &signalMapping, &segmentMapping)
	findSignalZero(data, &signalMapping, &segmentMapping)
	findSignalSix(data, &signalMapping, &segmentMapping)
	findSignalTwo(data, &signalMapping, &segmentMapping)
	findSignalFive(data, &signalMapping, &segmentMapping)

	return signalMapping, segmentMapping
}
func findSignalZero(data dataLine, signalMapping *map[int]string, segmentMapping *map[int]string) {
	(*signalMapping)[0] = sortString((*signalMapping)[1] + (*segmentMapping)[0] + (*segmentMapping)[1] + (*segmentMapping)[4] + (*segmentMapping)[6])

	difference := findSegmentsDifference((*signalMapping)[8], (*signalMapping)[0])
	runtimeAssert(len(difference) == 1)
	(*segmentMapping)[3] = difference[0]
}
func findSignalOne(data dataLine, signalMapping *map[int]string, segmentMapping *map[int]string) {
	oneSignal := getSignalsWithLength(data, 2)
	runtimeAssert(len(oneSignal) == 1)
	(*signalMapping)[1] = oneSignal[0]
}
func findSignalTwo(data dataLine, signalMapping *map[int]string, segmentMapping *map[int]string) {
	(*signalMapping)[2] = sortString((*segmentMapping)[0] + (*segmentMapping)[2] + (*segmentMapping)[3] + (*segmentMapping)[4] + (*segmentMapping)[6])
}
func findSignalThree(data dataLine, signalMapping *map[int]string, segmentMapping *map[int]string) {
	base := sortString((*signalMapping)[1] + (*segmentMapping)[0] + (*segmentMapping)[6])
	signalsWith5Segments := getSignalsWithLength(data, 5)
	runtimeAssert(len(signalsWith5Segments) == 3)

	for _, signal := range signalsWith5Segments {
		diff := findSegmentsDifference(base, signal)
		if len(diff) == 0 {
			runtimeAssert(len((*signalMapping)[3]) == 0)
			(*signalMapping)[3] = signal
		}
	}

	runtimeAssert(len((*signalMapping)[3]) != 0)

	almostThree := sortString((*signalMapping)[4] + (*segmentMapping)[0] + (*segmentMapping)[6])
	difference := findSegmentsDifference(almostThree, (*signalMapping)[3])
	runtimeAssert(len(difference) == 1)
	(*segmentMapping)[1] = difference[0]
}
func findSignalFour(data dataLine, signalMapping *map[int]string, segmentMapping *map[int]string) {
	fourSignal := getSignalsWithLength(data, 4)
	runtimeAssert(len(fourSignal) == 1)
	(*signalMapping)[4] = fourSignal[0]
}
func findSignalFive(data dataLine, signalMapping *map[int]string, segmentMapping *map[int]string) {
	signalsWith5Segments := getSignalsWithLength(data, 5)
	runtimeAssert(len(signalsWith5Segments) == 3)

	for _, signal := range signalsWith5Segments {
		if signal == (*signalMapping)[2] || signal == (*signalMapping)[3] {
			continue
		}

		runtimeAssert(len((*signalMapping)[5]) == 0)
		(*signalMapping)[5] = signal
	}

	runtimeAssert(len((*signalMapping)[5]) != 0)
}
func findSignalSix(data dataLine, signalMapping *map[int]string, segmentMapping *map[int]string) {
	signalsWith6Segments := getSignalsWithLength(data, 6)
	runtimeAssert(len(signalsWith6Segments) == 3)

	for _, signal := range signalsWith6Segments {
		if signal == (*signalMapping)[0] || signal == (*signalMapping)[9] {
			continue
		}

		runtimeAssert(len((*signalMapping)[6]) == 0)
		(*signalMapping)[6] = signal
	}

	runtimeAssert(len((*signalMapping)[6]) != 0)

	missingSegment := findSegmentsDifference((*signalMapping)[8], (*signalMapping)[6])
	runtimeAssert(len(missingSegment) == 1)
	(*segmentMapping)[2] = missingSegment[0]
}
func findSignalSeven(data dataLine, signalMapping *map[int]string, segmentMapping *map[int]string) {
	sevenSignal := getSignalsWithLength(data, 3)
	runtimeAssert(len(sevenSignal) == 1)
	(*signalMapping)[7] = sevenSignal[0]

	inSeven := findSegmentsDifference((*signalMapping)[7], (*signalMapping)[1])
	runtimeAssert(len(inSeven) == 1)
	(*segmentMapping)[0] = inSeven[0]
}
func findSignalEight(data dataLine, signalMapping *map[int]string, segmentMapping *map[int]string) {
	eightSignal := getSignalsWithLength(data, 7)
	runtimeAssert(len(eightSignal) == 1)
	(*signalMapping)[8] = eightSignal[0]
}
func findSignalNine(data dataLine, signalMapping *map[int]string, segmentMapping *map[int]string) {
	almostNineSignalString := (*signalMapping)[4] + (*segmentMapping)[0]
	sortString(almostNineSignalString)
	missingSegments := getMissingSegments(almostNineSignalString)
	runtimeAssert(len(missingSegments) == 2)

	candidate0 := sortString(almostNineSignalString + missingSegments[0])
	candidate1 := sortString(almostNineSignalString + missingSegments[1])
	if findString(candidate0, data.signalPatterns[:]) {
		runtimeAssert(!findString(candidate1, data.signalPatterns[:]))

		(*signalMapping)[9] = candidate0
		(*segmentMapping)[4] = missingSegments[1]
		(*segmentMapping)[6] = missingSegments[0]
	} else {
		runtimeAssert(findString(candidate1, data.signalPatterns[:]))
		(*signalMapping)[9] = candidate1
		(*segmentMapping)[6] = missingSegments[1]
		(*segmentMapping)[4] = missingSegments[0]
	}
}
func findSegmentsDifference(original, target string) []string {
	result := []string{}

	originalRunes := []rune(original)
	targetRunes := []rune(target)

	for _, originalRune := range originalRunes {
		if findRune(originalRune, targetRunes) {
			continue
		}

		result = append(result, string(originalRune))
	}

	return result
}
func findRune(rune rune, targetRunes []rune) bool {
	for _, targetRune := range targetRunes {
		if rune == targetRune {
			return true
		}
	}
	return false
}
func findString(s string, targetStrings []string) bool {
	for _, targetString := range targetStrings {
		if s == targetString {
			return true
		}
	}
	return false
}
func getSignalsWithLength(data dataLine, length int) []string {
	result := []string{}

	for _, item := range data.signalPatterns {
		if len(item) == length {
			result = append(result, item)
		}
	}

	return result
}
func getMissingSegments(signal string) []string {
	return findSegmentsDifference("abcdefg", signal)
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
		result = append(result, parseLine(scanner.Text()))
	}
	return result
}
func parseLine(data string) dataLine {
	fields := strings.Fields(data)

	return dataLine{
		signalPatterns: [10]string{
			sortString(fields[0]),
			sortString(fields[1]),
			sortString(fields[2]),
			sortString(fields[3]),
			sortString(fields[4]),
			sortString(fields[5]),
			sortString(fields[6]),
			sortString(fields[7]),
			sortString(fields[8]),
			sortString(fields[9]),
		},
		output: [4]string{
			sortString(fields[11]),
			sortString(fields[12]),
			sortString(fields[13]),
			sortString(fields[14]),
		},
	}
}
