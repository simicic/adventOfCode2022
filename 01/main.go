package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	maxCalories := MaxCalories("../input_test.txt")
	fmt.Println("Max calories: ", maxCalories)

	maxTopThreeCalories := MaxTopThreeCalories("../input_test.txt")
	fmt.Println("Max top three calories: ", maxTopThreeCalories)
}

func MaxTopThreeCalories(fileName string) int {
	caloriesPerElf := make(map[int]int)
	caloriesSum := 0
	elfCount := 1

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		calories, error := strconv.Atoi(line)
		if error != nil {
			// newline - finalize calories for an elf
			caloriesPerElf[elfCount] = caloriesSum

			caloriesSum = 0
			elfCount++
		}
		caloriesSum += calories
	}

	caloriesPerElf[elfCount] = caloriesSum

	slice := []int{len(caloriesPerElf)}
	for _, value := range caloriesPerElf {
		slice = append(slice, value)
	}

	sort.Ints(slice)
	n := len(slice)

	return slice[n-1] + slice[n-2] + slice[n-3]
}

func MaxCalories(fileName string) int {
	maxCalories := 0
	caloriesPerElf := 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		calories, error := strconv.Atoi(line)
		if error != nil {
			// newline - finalize calories for an elf
			if caloriesPerElf > maxCalories {
				maxCalories = caloriesPerElf
			}
			caloriesPerElf = 0
		}
		caloriesPerElf += calories
	}

	return maxCalories
}
