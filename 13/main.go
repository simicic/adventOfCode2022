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
	correctlyOrderedIndicesCount := CorrectlyOrderedIndicesCount("13/input_test.txt")
	fmt.Println("Sum of correctly ordered indices: ", correctlyOrderedIndicesCount)
}

func CorrectlyOrderedIndicesCount(fileName string) int {
	sum := 1
	var correctlyOrdered []int

	inputData := readInputData(fileName)

	for i, v := range inputData {
		fmt.Println("Pair ", i+1)
		if validateOrder(v[0], v[1]) == true {
			fmt.Println("Outcome: TRUE")
			correctlyOrdered = append(correctlyOrdered, i)
		} else {
			fmt.Println("Outcome: FALSE")
		}
	}

	fmt.Println(correctlyOrdered)

	return sum
}

func validateOrder(line1 string, line2 string) bool {
	fmt.Println("Comparing: ", line1, " with ", line2)

	outcome := false

	if strings.HasPrefix(line1, "[") == true || strings.HasPrefix(line2, "[") == true {
		line1, line2 = refactorBraces(line1, line2)
		validateOrder(line1, line2)
	}

	line1Elem := strings.Split(line1, ",")
	line2Elem := strings.Split(line2, ",")

	if len(line1Elem) == 0 {
		return false
	}

	for i, l1v := range line1Elem {
		v1, _ := strconv.Atoi(l1v)

		if i >= len(line2Elem) {
			return false
		}

		v2, _ := strconv.Atoi(line2Elem[i])

		if v1 < v2 {
			return true
		} else if v1 > v2 {
			return false
		}
	}

	return outcome
}

func refactorBraces(line1 string, line2 string) (string, string) {
	if strings.HasPrefix(line1, "[") == true {
		line1 = replace(line1)
	} else {
		line1 = fmt.Sprintf("[%v]", line1)
	}

	if strings.HasPrefix(line2, "[") == true {
		line2 = replace(line2)
	} else {
		line2 = fmt.Sprintf("[%v]", line2)
	}

	return line1, line2
}

func replace(line string) string {
	line = strings.TrimLeft(line, "[")
	line = strings.TrimRight(line, "]")
	return line
}

func readInputData(fileName string) [][]string {
	//[one pair, other pair], []
	var inputData [][]string

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()

		pair := []string{line1, line2}
		inputData = append(inputData, pair)

		scanner.Scan()
	}

	return inputData
}
