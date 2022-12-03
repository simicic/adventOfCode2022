package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	priority := Priority("03/input.txt")
	fmt.Println("Priority: ", priority)

	groupPriority := GroupPriority("03/input.txt")
	fmt.Println("Group priority: ", groupPriority)
}

func Priority(fileName string) int {
	priority := 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		middle := len(line) / 2

		firstHalf := line[:middle]
		secondHalf := line[middle:]

	out:
		for _, n := range firstHalf {
			for _, m := range secondHalf {
				if m == n {
					priority += assignedPriority(n)
					break out
				}
			}
		}
	}

	return priority
}

func GroupPriority(fileName string) int {
	priority := 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		first := scanner.Text()
		scanner.Scan()
		second := scanner.Text()
		scanner.Scan()
		third := scanner.Text()

		common := findCommon(first, second, third)
		priority += assignedPriority(common)
	}

	return priority
}

func assignedPriority(letter rune) int {
	if letter <= 90 {
		// translate capital
		return int(letter) - 38
	} else {
		// translate lower case
		return int(letter) - 96
	}
}

func findCommon(first string, second string, third string) rune {
	for _, i := range first {
		for _, j := range second {
			for _, k := range third {
				if i == j && j == k {
					return i
				}
			}
		}
	}
	panic("Ouch!")
}
