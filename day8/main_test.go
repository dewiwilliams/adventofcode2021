package main

import (
	"testing"
)

func TestFindValues(t *testing.T) {
	testSignal := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
	parsed := parseLine(testSignal)

	signalMapping, _ := getMappings(parsed)

	expectedMapping := map[int]string{
		0: sortString("cagedb"),
		1: sortString("ab"),
		2: sortString("gcdfa"),
		3: sortString("fbcad"),
		4: sortString("eafb"),
		5: sortString("cdfbe"),
		6: sortString("cdfgeb"),
		7: sortString("dab"),
		8: sortString("acedgfb"),
		9: sortString("cefabd"),
	}

	for index, item := range expectedMapping {
		if signalMapping[index] != item {
			t.Errorf("Signal failed, %d: %s vs %s", index, signalMapping[index], item)
		}
	}
}
