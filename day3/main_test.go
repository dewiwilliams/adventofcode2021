package main

import (
	"testing"
)

var testData = []int64{
	0b00100,
	0b11110,
	0b10110,
	0b10111,
	0b10101,
	0b01111,
	0b00111,
	0b11100,
	0b10000,
	0b11001,
	0b00010,
	0b01010,
}

func TestGetDataSetScore(t *testing.T) {
	expectedScores := []int{-2, 2, 4, -2, 2}

	for i := 0; i < 5; i++ {
		score := getDataSetScore(testData, i)

		if score != expectedScores[i] {
			t.Errorf("Score does not match expected score[%d]: %d vs %d", i, score, expectedScores[i])
		}
	}
}

func TestFilterDataSet1(t *testing.T) {
	var expectedResult = []int64{
		0b11110,
		0b10110,
		0b10111,
		0b01111,
		0b00111,
		0b00010,
		0b01010,
	}

	result := filterDataSet(testData, 1, true)

	if !equal(result, expectedResult) {
		t.Errorf("Expected filter does not match")
	}
}
func TestFilterDataSet2(t *testing.T) {
	var expectedResult = []int64{
		0b00100,
		0b01111,
		0b00111,
		0b00010,
		0b01010,
	}

	result := filterDataSet(testData, 4, false)

	if !equal(result, expectedResult) {
		t.Errorf("Expected filter does not match")
	}
}

func TestOxygenRating(t *testing.T) {
	oxygenRating := getOxygenRating(testData, 5)

	if oxygenRating != 23 {
		t.Errorf("Oxygen rating failed: %d vs %d", oxygenRating, 23)
	}
}

func TestScrubberRating(t *testing.T) {
	scrubberRating := getScrubberRating(testData, 5)

	if scrubberRating != 10 {
		t.Errorf("Scrubber rating failed: %d vs %d", scrubberRating, 10)
	}
}

func equal(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
