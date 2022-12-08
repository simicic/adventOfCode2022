package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	visibleTrees := VisibleTrees("08/input.txt")
	fmt.Println("Visible trees: ", visibleTrees)
}

func VisibleTrees(fileName string) int {
	visibleTrees := 0

	inputData := readInputData(fileName)
	visibleTrees += visibleTreesOnEdge(inputData)
	visibleTrees += visibleTreesInside(inputData)

	return visibleTrees
}

func visibleTreesInside(inputData [][]int) int {
	count := 0

	iMax := len(inputData) - 1
	jMax := len(inputData) - 1

	for i := 1; i < iMax; i++ {
		for j := 1; j < jMax; j++ {
			if isVisible(inputData, i, j, iMax, jMax) == true {
				count++
			}
		}
	}

	return count
}

func isVisible(inputData [][]int, x int, y int, iMax int, jMax int) bool {
	currentTree := inputData[x][y]

	allHigherLeft := false
	allHigherRight := false
	allHigherUp := false
	allHigherDown := false

	// left
	for j := y - 1; j >= 0; j-- {
		if inputData[x][j] >= currentTree {
			allHigherLeft = true
			break
		}
	}

	// right
	for j := y + 1; j < iMax+1; j++ {
		if inputData[x][j] >= currentTree {
			allHigherRight = true
			break
		}
	}

	// up
	for i := x - 1; i >= 0; i-- {
		if inputData[i][y] >= currentTree {
			allHigherUp = true
			break
		}
	}

	// down
	for i := x + 1; i < jMax+1; i++ {
		if inputData[i][y] >= currentTree {
			allHigherDown = true
			break
		}
	}

	if allHigherLeft == false ||
		allHigherRight == false ||
		allHigherUp == false ||
		allHigherDown == false {
		return true
	} else {
		return false
	}
}

func visibleTreesOnEdge(inputData [][]int) int {
	count := 0
	count += len(inputData) * 2
	count += (len(inputData[0]) - 2) * 2
	return count
}

func readInputData(fileName string) [][]int {
	var inputData [][]int
	i := 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var lineArray []int

		for j := 0; j < len(line); j++ {
			num, _ := strconv.Atoi(string(line[j]))
			lineArray = append(lineArray, num)
		}

		inputData = append(inputData, lineArray)
		i++
	}

	return inputData
}
