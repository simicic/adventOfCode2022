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
