package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	if getPart1Score("{([(<{}[<>[]}>{[]{[(<()>") != 1197 {
		t.Errorf("Wrong score 1: %d vs %d", getPart1Score("{([(<{}[<>[]}>{[]{[(<()>"), 1197)
	}

	if getPart1Score("[[<[([]))<([[{}[[()]]]") != 3 {
		t.Errorf("Wrong score 2: %d vs %d", getPart1Score("[[<[([]))<([[{}[[()]]]"), 3)
	}

	if getPart1Score("[{[{({}]{}}([{[{{{}}([]") != 57 {
		t.Errorf("Wrong score 3: %d vs %d", getPart1Score("[{[{({}]{}}([{[{{{}}([]"), 57)
	}

	if getPart1Score("[<(<(<(<{}))><([]([]()") != 3 {
		t.Errorf("Wrong score 4: %d vs %d", getPart1Score("[<(<(<(<{}))><([]([]()"), 3)
	}

	if getPart1Score("<{([([[(<>()){}]>(<<{{") != 25137 {
		t.Errorf("Wrong score 5: %d vs %d", getPart1Score("<{([([[(<>()){}]>(<<{{"), 25137)
	}
}
