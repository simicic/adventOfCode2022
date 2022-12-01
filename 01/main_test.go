package main

import "testing"

var maxCaloriesTests = []struct {
	input    string
	expected int // expected result
}{
	{"input_test.txt", 24000},
}

func TestMaxCalories(t *testing.T) {
	for _, val := range maxCaloriesTests {
		got := MaxCalories(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}

var maxTopThreeTests = []struct {
	input    string
	expected int // expected result
}{
	{"input_test.txt", 45000},
}

func TestMaxTopThreeCalories(t *testing.T) {
	for _, val := range maxTopThreeTests {
		got := MaxTopThreeCalories(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}
