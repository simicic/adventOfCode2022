package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	topCratesCode := TopCratesCode("05/input.txt")
	fmt.Println("Top crates code: ", topCratesCode)

	topCrateMover9001Code := TopCrateMover9001("05/input.txt")
	fmt.Println("Top crates code with CrateMover 9001: ", topCrateMover9001Code)
}

func TopCratesCode(fileName string) string {
	code := ""
	stacksInfo := make([]string, 10)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stacksInfo = append(stacksInfo, line)

		if len(line) == 0 {
			stacksOfCrates := formatInputData(stacksInfo)

			// process commands
			for scanner.Scan() {
				commandLine := scanner.Text()

				commandLineRegexp := regexp.MustCompile(`move (.+) from (.) to (.)`)
				moves := commandLineRegexp.FindStringSubmatch(commandLine)

				countToMove, _ := strconv.Atoi(moves[1])
				from := convertSource(moves[2])
				to := convertSource(moves[3])

				for i := 0; i < countToMove; i++ {
					elementToMove := stacksOfCrates[from][len(stacksOfCrates[from])-1]

					stacksOfCrates[from] = stacksOfCrates[from][:len(stacksOfCrates[from])-1]
					stacksOfCrates[to] = append(stacksOfCrates[to], elementToMove)
				}
			}

			for _, v := range stacksOfCrates {
				code = code + v[len(v)-1]
			}

			break
		}
	}

	return code
}

func TopCrateMover9001(fileName string) string {
	code := ""
	stacksInfo := make([]string, 10)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stacksInfo = append(stacksInfo, line)

		if len(line) == 0 {
			stacksOfCrates := formatInputData(stacksInfo)
			// process commands
			for scanner.Scan() {
				commandLine := scanner.Text()

				commandLineRegexp := regexp.MustCompile(`move (.+) from (.) to (.)`)
				moves := commandLineRegexp.FindStringSubmatch(commandLine)

				countToMove := convertSource(moves[1])
				from := convertSource(moves[2])
				to := convertSource(moves[3])

				elementsToMove := stacksOfCrates[from][len(stacksOfCrates[from])-countToMove-1:]

				stacksOfCrates[from] = stacksOfCrates[from][:len(stacksOfCrates[from])-countToMove-1]
				stacksOfCrates[to] = append(stacksOfCrates[to], elementsToMove...)
			}

			for _, v := range stacksOfCrates {
				code = code + v[len(v)-1]
			}

			break
		}
	}

	return code
}

func convertSource(input string) int {
	num, _ := strconv.Atoi(input)
	return num - 1
}
func formatInputData(stacksInfo []string) [][]string {
	stacksCountLine := strings.Split(stacksInfo[len(stacksInfo)-2], " ")
	length, _ := strconv.Atoi(stacksCountLine[len(stacksCountLine)-1])

	stacksOfCrates := make([][]string, length)
	k := 0

	for i := len(stacksInfo) - 3; i > 0; i-- {
		k = 0
		for j := 1; j < len(stacksInfo[i]); j = j + 4 {
			crateCode := string(stacksInfo[i][j])

			if crateCode != " " {
				stacksOfCrates[k] = append(stacksOfCrates[k], crateCode)
			}

			k++
		}
	}

	return stacksOfCrates
}
