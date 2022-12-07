package main

import "testing"

var totalSizeTest = []struct {
	input    string
	expected int // expected result
}{
	{"input_test.txt", 95437},
}

func TestTotalSize(t *testing.T) {
	for _, val := range totalSizeTest {
		got := TotalSize(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}
