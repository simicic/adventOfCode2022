package main

import "testing"

var topCratesCodeTest = []struct {
	input    string
	expected string // expected result
}{
	{"input_test.txt", "CMZ"},
}

func TestTopCratesCode(t *testing.T) {
	for _, val := range topCratesCodeTest {
		got := TopCratesCode(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}
