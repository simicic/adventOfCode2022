package main

import "testing"

var validateOrderTest = []struct {
	line1    string
	line2    string
	expected bool
}{
	{line1: "[1,1,3,1,1]", line2: "[1,1,5,1,1]", expected: true},
	{line1: "[[1],[2,3,4]]", line2: "[[1],4]", expected: true},
	{line1: "[9]", line2: "[[8,7,6]]", expected: false},
	{line1: "[[4,4],4,4]", line2: "[[4,4],4,4,4]", expected: true},
	{line1: "[7,7,7,7]", line2: "[7,7,7]", expected: false},
	{line1: "[]", line2: "[3]", expected: true},
	{line1: "[[[]]]", line2: "[[]]", expected: false},
	{line1: "[1,[2,[3,[4,[5,6,7]]]],8,9]", line2: "[1,[2,[3,[4,[5,6,0]]]],8,9]", expected: false},
}

func TestValidateOrder(t *testing.T) {
	for _, val := range validateOrderTest {
		got := ValidateOrder(val.line1, val.line2)

		if got != val.expected {
			t.Error("Comparing ", val.line1, " and ", val.line2, " was not correct, it was %v: ", got)
		}
	}
}
