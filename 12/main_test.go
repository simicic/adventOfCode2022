package main

import "testing"

var shortestPathTest = []struct {
	input    string
	expected int // expected result
}{
	{"input_test.txt", 31},
}

func TestShortestPath(t *testing.T) {
	for _, val := range shortestPathTest {
		got := ShortestPath(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}
