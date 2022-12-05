package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	duplicatePairs := DuplicatePairs("04/input.txt")
	fmt.Println("Duplicate pairs: ", duplicatePairs)

	duplicateOverallPairs := DuplicateOverallPairs("04/input.txt")
	fmt.Println("Duplicate pairs overall: ", duplicateOverallPairs)
}

func DuplicatePairs(fileName string) int {
	pairs := 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")

		startFirst, endFirst := startAndEnd(line[0])
		startSecond, endSecond := startAndEnd(line[1])

		if (startFirst <= startSecond && endFirst >= endSecond) || startSecond <= startFirst && endSecond >= endFirst {
			pairs += 1
		}
	}

	return pairs
}

func DuplicateOverallPairs(fileName string) int {
	pairs := 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")

		startFirst, endFirst := startAndEnd(line[0])
		startSecond, endSecond := startAndEnd(line[1])

		if startSecond <= endFirst && endSecond >= startFirst {
			pairs += 1
		}
	}

	return pairs
}

func startAndEnd(line string) (int, int) {
	startAndEnd := strings.Split(line, "-")

	start, _ := strconv.Atoi(startAndEnd[0])
	end, _ := strconv.Atoi(startAndEnd[1])

	return start, end
}
