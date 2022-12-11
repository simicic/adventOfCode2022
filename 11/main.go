package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items               []int
	inspectedItemsCount int
	operation           string
	testDivisionBy      int
	divisionTrue        int
	divisionFalse       int
}

func main() {
	monkeyBusinessLevel := MonkeyBusinessLevel("11/input_test.txt")
	fmt.Println("Level of monkey business: ", monkeyBusinessLevel)
}

func MonkeyBusinessLevel(fileName string) int {
	monkeyBusinessLevel := 1
	roundCount := 1

	monkeyBusiness := readInputData(fileName)
	roundsLimit := 20

	for roundCount <= roundsLimit {
		for monkeyIndex := 0; monkeyIndex < len(monkeyBusiness); monkeyIndex++ {
			monkey := monkeyBusiness[monkeyIndex]

			for _, item := range monkey.items {
				newValue := processOperation(monkey.operation, item)
				newValue = newValue / 3

				var otherMonkey Monkey
				var otherMonkeyIndex int

				if newValue%monkey.testDivisionBy == 0 {
					otherMonkeyIndex = monkey.divisionTrue
					otherMonkey = monkeyBusiness[otherMonkeyIndex]
				} else {
					otherMonkeyIndex = monkey.divisionFalse
					otherMonkey = monkeyBusiness[otherMonkeyIndex]
				}

				otherMonkey.items = append(otherMonkey.items, newValue)

				delete(monkeyBusiness, otherMonkeyIndex)
				monkeyBusiness[otherMonkeyIndex] = otherMonkey

				(&monkey).inspectedItemsCount = (&monkey).inspectedItemsCount + 1
			}

			(&monkey).items = []int{}

			delete(monkeyBusiness, monkeyIndex)
			monkeyBusiness[monkeyIndex] = monkey
		}

		roundCount++
		if roundCount > roundsLimit {
			break
		}
	}

	var levels []int
	for _, monkey := range monkeyBusiness {
		levels = append(levels, monkey.inspectedItemsCount)
	}
	sort.Ints(levels)

	monkeyBusinessLevel = levels[len(levels)-1] * levels[len(levels)-2]

	return monkeyBusinessLevel
}

func printCountsPerMonkey(monkeyBusiness map[int]Monkey) {
	for i, v := range monkeyBusiness {
		fmt.Println("Monkey: ", i, " : ", v.items)
	}
}

func processOperation(operation string, item int) int {
	operationRegexp := regexp.MustCompile(`old (.) (.+)`)
	operationSlice := operationRegexp.FindStringSubmatch(operation)

	newValue := 0
	by := 0

	if operationSlice[2] == "old" {
		by = item
	} else {
		by, _ = strconv.Atoi(operationSlice[2])
	}

	if operationSlice[1] == "*" {
		newValue = item * by
	} else {
		newValue = item + by
	}

	return newValue
}

func readInputData(fileName string) map[int]Monkey {
	inputData := make(map[int]Monkey)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Monkey ") {
			// Monkey 0:
			monkeyRegexp := regexp.MustCompile(`Monkey (.+):`)
			code := monkeyRegexp.FindStringSubmatch(line)[1]
			codeInt, _ := strconv.Atoi(code)

			// Starting items: 79, 98
			scanner.Scan()
			startingItems := scanner.Text()
			startingItems = strings.Replace(startingItems, " Starting items: ", "", 1)
			itemsSlice := strings.Split(startingItems, ",")

			var items []int
			for _, v := range itemsSlice {
				vInt, _ := strconv.Atoi(strings.Trim(v, " "))
				items = append(items, vInt)
			}

			// Operation: new = old * 19
			scanner.Scan()
			operation := scanner.Text()
			operation = strings.Replace(operation, "  Operation: new = ", "", 1)

			// Test: divisible by 23
			scanner.Scan()
			divisibleBy := scanner.Text()
			divisibleBy = strings.Replace(divisibleBy, " Test: divisible by ", "", 1)
			divisibleByInt, _ := strconv.Atoi(strings.Trim(divisibleBy, " "))

			// If true: throw to monkey 2
			scanner.Scan()
			ifTrue := scanner.Text()
			ifTrue = strings.Replace(ifTrue, "    If true: throw to monkey ", "", 1)
			ifTrueInt, _ := strconv.Atoi(strings.Trim(ifTrue, " "))

			// If false: throw to monkey 3
			scanner.Scan()
			ifFalse := scanner.Text()
			ifFalse = strings.Replace(ifFalse, "    If false: throw to monkey ", "", 1)
			ifFalseInt, _ := strconv.Atoi(strings.Trim(ifFalse, " "))

			monkey := Monkey{
				items:          items,
				operation:      operation,
				testDivisionBy: divisibleByInt,
				divisionTrue:   ifTrueInt,
				divisionFalse:  ifFalseInt,
			}

			inputData[codeInt] = monkey
		}
	}

	return inputData
}
