package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	visibleTrees := VisibleTrees("08/input_test.txt")
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

	var scenicScores []int

	for i := 1; i < iMax; i++ {
		for j := 1; j < jMax; j++ {
			visible, scenicScore := isVisible(inputData, i, j, iMax, jMax)
			if visible == true {
				count++
			}
			scenicScores = append(scenicScores, scenicScore)
		}
	}

	// max
	fmt.Println("Scenic scores: ", scenicScores)

	return count
}

func isVisible(inputData [][]int, x int, y int, iMax int, jMax int) (bool, int) {
	currentTree := inputData[x][y]
	scenicScore := 0

	allHigherLeft := false
	allHigherRight := false
	allHigherUp := false
	allHigherDown := false

	scenicScoreLeft := 0
	scenicScorerRight := 0
	scenicScoreUp := 0
	scenicScoreDown := 0

	// left
	for j := y - 1; j >= 0; j-- {
		if inputData[x][j] >= currentTree {
			allHigherLeft = true
			break
		} else {
			scenicScoreLeft++
		}
	}

	// right
	for j := y + 1; j < iMax+1; j++ {
		if inputData[x][j] >= currentTree {
			allHigherRight = true
			break
		} else {
			scenicScorerRight++
		}
	}

	// up
	for i := x - 1; i >= 0; i-- {
		if inputData[i][y] >= currentTree {
			allHigherUp = true
			break
		} else {
			scenicScoreUp++
		}
	}

	// down
	for i := x + 1; i < jMax+1; i++ {
		if inputData[i][y] >= currentTree {
			allHigherDown = true
			break
		} else {
			scenicScoreDown++
		}
	}

	scenicScore = scenicScoreLeft * scenicScorerRight * scenicScoreUp * scenicScoreDown

	if allHigherLeft == false ||
		allHigherRight == false ||
		allHigherUp == false ||
		allHigherDown == false {
		return true, scenicScore
	} else {
		return false, scenicScore
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
