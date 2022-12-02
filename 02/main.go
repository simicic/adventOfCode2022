package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Action struct {
	secretCode string
	points     int
	canWin     string
	canLoose   string
}

func main() {
	myScore := MyScore("02/input.txt")
	fmt.Println("My score: ", myScore)

	myScorePartTwo := MyScorePartTwo("02/input.txt")
	fmt.Println("My score: ", myScorePartTwo)
}

func MyScore(fileName string) int {
	// A = Rock => X (1)
	// B = Paper => Y (2)
	// C = Scissors => Z (3)

	rock := Action{secretCode: "X", points: 1, canWin: "C"}
	paper := Action{secretCode: "Y", points: 2, canWin: "A"}
	scissors := Action{secretCode: "Z", points: 3, canWin: "B"}

	actionValuations := make(map[string]Action, 3)

	actionValuations["A"] = rock
	actionValuations["B"] = paper
	actionValuations["C"] = scissors

	codes := map[string]string{"X": "A", "Y": "B", "Z": "C"}
	myScore := 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		if line[0] == codes[line[1]] {
			// 3 = draw
			myScore += actionValuations[codes[line[1]]].points + 3
		} else {
			if codes[line[1]] == actionValuations[line[0]].canWin {
				// 0 = lost
				myScore += actionValuations[codes[line[1]]].points
			} else {
				// 6 = won
				myScore += actionValuations[codes[line[1]]].points + 6
			}
		}
	}

	return myScore
}

func MyScorePartTwo(fileName string) int {
	// A = Rock => X (1)
	// B = Paper => Y (2)
	// C = Scissors => Z (3)

	rock := Action{secretCode: "X", points: 1, canWin: "C", canLoose: "B"}
	paper := Action{secretCode: "Y", points: 2, canWin: "A", canLoose: "C"}
	scissors := Action{secretCode: "Z", points: 3, canWin: "B", canLoose: "A"}

	actionValuations := make(map[string]Action, 3)

	actionValuations["A"] = rock
	actionValuations["B"] = paper
	actionValuations["C"] = scissors

	myScore := 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		var shouldPlay string

		if line[1] == "X" {
			// X means you need to loose
			shouldPlay = actionValuations[line[0]].canWin
			myScore += actionValuations[shouldPlay].points
		} else if line[1] == "Y" {
			// Y means you need to end the round in a draw
			myScore += actionValuations[line[0]].points + 3
		} else {
			shouldPlay = actionValuations[line[0]].canLoose
			myScore += actionValuations[shouldPlay].points + 6
		}
	}

	return myScore
}
