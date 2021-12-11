package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testData := []int{
		2, 1, 9, 9, 9, 4, 3, 2, 1, 0,
		3, 9, 8, 7, 8, 9, 4, 9, 2, 1,
		9, 8, 5, 6, 7, 8, 9, 8, 9, 2,
		8, 7, 6, 7, 8, 9, 6, 7, 8, 9,
		9, 8, 9, 9, 9, 6, 5, 6, 7, 8,
	}

	risklevel := getRiskLevel(testData, 10)
	if risklevel != 15 {
		t.Errorf("Wrong risk level: %d vs %d", risklevel, 15)
	}
}

func TestPart2(t *testing.T) {
	testData := []int{
		2, 1, 9, 9, 9, 4, 3, 2, 1, 0,
		3, 9, 8, 7, 8, 9, 4, 9, 2, 1,
		9, 8, 5, 6, 7, 8, 9, 8, 9, 2,
		8, 7, 6, 7, 8, 9, 6, 7, 8, 9,
		9, 8, 9, 9, 9, 6, 5, 6, 7, 8,
	}

	basin1Size := getBasinSize(testData, 10, 1)
	if basin1Size != 3 {
		t.Errorf("Wrong basin 1 size: %d vs %d", basin1Size, 3)
	}

	basin2Size := getBasinSize(testData, 10, 9)
	if basin2Size != 9 {
		t.Errorf("Wrong basin 2 size: %d vs %d", basin2Size, 9)
	}

	basin3Size := getBasinSize(testData, 10, 22)
	if basin3Size != 14 {
		t.Errorf("Wrong basin 3 size: %d vs %d", basin3Size, 14)
	}

	basin4Size := getBasinSize(testData, 10, 46)
	if basin4Size != 9 {
		t.Errorf("Wrong basin 4 size: %d vs %d", basin4Size, 9)
	}
}
