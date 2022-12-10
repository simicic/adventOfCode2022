package main

import "testing"

var visitedCountTest = []struct {
	input    string
	expected int // expected result
}{
	{"input_test.txt", 13},
}

func TestVisitedCount(t *testing.T) {
	for _, val := range visitedCountTest {
		got := VisitedCount(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}

var visitedCountLongTest = []struct {
	input    string
	expected int // expected result
}{
	{"input_test_01.txt", 36},
}

func TestVisitedCountLong(t *testing.T) {
	for _, val := range visitedCountLongTest {
		got := VisitedCount(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}
