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
	signalStrengthTotal := SignalStrengthTotal("10/input.txt")
	fmt.Println("Signal strength total: ", signalStrengthTotal)
}

func SignalStrengthTotal(fileName string) int {
	signalStrengthTotal := 0
	iterationValues := map[int]int{0: 1}

	crtScreen := initializeCrtScreen()
	inputData := readInputData(fileName)

	cycleCount := 1
	inputDataIndex := 0

	for {
		if cycleCount == 20 || cycleCount == 60 || cycleCount == 100 || cycleCount == 140 || cycleCount == 180 || cycleCount == 220 {
			signalStrengthTotal += iterationValues[cycleCount-1] * cycleCount
		}

		if strings.HasPrefix(inputData[inputDataIndex], "addx") {
			command := strings.Replace(inputData[inputDataIndex], "addx ", "", 1)
			commandStep := strings.Split(command, " ")

			if len(commandStep) > 1 {
				// action because it is second time touching the command
				step, _ := strconv.Atoi(commandStep[0])
				iterationValues[cycleCount] = iterationValues[cycleCount-1] + step
				inputDataIndex++
			} else {
				inputData[inputDataIndex] = fmt.Sprintf("%v 1", inputData[inputDataIndex])
				iterationValues[cycleCount] = iterationValues[cycleCount-1]
			}
		} else {
			// does nothing
			iterationValues[cycleCount] = iterationValues[cycleCount-1]
			inputDataIndex++
		}

		cycleCount++

		if inputDataIndex >= len(inputData) {
			break
		}
	}

	for i := 0; i < len(iterationValues)-1; i++ {
		column := i % 40
		row := i / 40

		x := iterationValues[i]
		distance := column - x

		if distance >= -1 && distance <= 1 {
			crtScreen[row][column] = "#"
		} else {
			crtScreen[row][column] = "."
		}
	}

	printCrtScreen(crtScreen)
	// PZGPKPEB

	return signalStrengthTotal
}

func printCrtScreen(crtScreen [][]string) {
	for _, v := range crtScreen {
		for _, item := range v {
			fmt.Print(item)
		}
		fmt.Print("\n")
	}
}

func initializeCrtScreen() [][]string {
	crtScreen := make([][]string, 6)

	for i := 0; i < 6; i++ {
		crtLine := make([]string, 40)

		for j := 0; j < 40; j++ {
			crtLine[j] = "."
		}

		crtScreen[i] = crtLine
	}

	return crtScreen
}

func readInputData(fileName string) []string {
	var inputData []string

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputData = append(inputData, line)
	}

	return inputData
}
