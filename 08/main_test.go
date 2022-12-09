package main

import "testing"

var visibleTreesTest = []struct {
	input    string
	expected int // expected result
}{
	{"input_test.txt", 21},
}

func TestVisibleTrees(t *testing.T) {
	for _, val := range visibleTreesTest {
		got := VisibleTrees(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}
