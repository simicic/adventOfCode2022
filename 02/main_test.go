package main

import "testing"

var myScoreTest = []struct {
	input    string
	expected int // expected result
}{
	{"input_test.txt", 15},
}

func TestMyScore(t *testing.T) {
	for _, val := range myScoreTest {
		got := MyScore(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}

var MyScorePartTwoTest = []struct {
	input    string
	expected int // expected result
}{
	{"input_test.txt", 12},
}

func TestMyScorePartTwo(t *testing.T) {
	for _, val := range MyScorePartTwoTest {
		got := MyScorePartTwo(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}
