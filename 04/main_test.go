package main

import "testing"

var duplicatePairsTest = []struct {
	input    string
	expected int // expected result
}{
	{"input_test.txt", 2},
}

func TestDuplicatePairs(t *testing.T) {
	for _, val := range duplicatePairsTest {
		got := DuplicatePairs(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}

var duplicateOverallPairsTest = []struct {
	input    string
	expected int // expected result
}{
	{"input_test.txt", 6},
}

func TestDuplicateOverallPairs(t *testing.T) {
	for _, val := range duplicateOverallPairsTest {
		got := DuplicateOverallPairs(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}
