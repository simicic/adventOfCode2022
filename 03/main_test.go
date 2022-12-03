package main

import "testing"

var priorityTest = []struct {
	input    string
	expected int // expected result
}{
	{"input_test.txt", 157},
}

func TestPriority(t *testing.T) {
	for _, val := range priorityTest {
		got := Priority(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}

var groupPriorityTest = []struct {
	input    string
	expected int // expected result
}{
	{"input_test.txt", 70},
}

func TestGroupPriority(t *testing.T) {
	for _, val := range groupPriorityTest {
		got := GroupPriority(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}
