package main

import "testing"

var topPacketStartIndex = []struct {
	input    string
	expected int // expected result
}{
	{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
	{"nppdvjthqldpwncqszvftbrmjlhg", 6},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
}

func TestPacketStartIndex(t *testing.T) {
	for _, val := range topPacketStartIndex {
		got := PacketStartIndex(val.input)

		if got != val.expected {
			t.Error("Nah, it was: ", got)
		}
	}
}

var topMessageStartIndex = []struct {
	input    string
	expected int // expected result
}{
	{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
	{"nppdvjthqldpwncqszvftbrmjlhg", 23},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
}

func TestMessageStartIndex(t *testing.T) {
	for _, val := range topMessageStartIndex {
		got := MessageStartIndex(val.input)

		if got != val.expected {
			t.Error("Nah, expected ", val.expected, "got: ", got)
		}
	}
}
