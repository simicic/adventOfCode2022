package main

import "testing"

var monkeyBusinessLevelTest = []struct {
	input    string
	expected int // expected result
}{
	{"input_test.txt", 10605},
}

func TestMonkeyBusinessLevel(t *testing.T) {
	for _, val := range monkeyBusinessLevelTest {
		got := MonkeyBusinessLevel(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}
