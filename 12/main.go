package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coordinate struct {
	x int
	y int
}

type Node struct {
	coordinate        Coordinate
	value             int
	distanceFromStart int
	visited           bool
}

type GridInfo struct {
	startNodeAt Coordinate
	endNodeAt   Coordinate
	sizeX       int
	sizeY       int
}

func main() {
	shortestPath := ShortestPath("12/input_test.txt")
	fmt.Println("Shortest path: ", shortestPath)
}

func ShortestPath(fileName string) int {
	gridInfo := getGridInfo(fileName)
	gridData := readInputData(fileName, gridInfo)

	startNode := gridData[gridInfo.startNodeAt.toString()]
	startNode.visited = true
	delete(gridData, startNode.coordinate.toString())
	gridData[startNode.coordinate.toString()] = startNode

	gridData = calculateDistances(gridData, gridInfo, startNode)

	printGridData(gridData, gridInfo)

	return 0
}

func printGridData(gridData map[string]Node, info GridInfo) {
	for y := 0; y < info.sizeY; y++ {
		var line []string
		for x := 0; x < info.sizeX; x++ {
			coordinate := Coordinate{x, y}.toString()
			v := "."
			if gridData[coordinate].visited == true {
				v = ">"
			}
			line = append(line, string(v))
		}
		fmt.Println(line)
	}
}

func calculateDistances(gridData map[string]Node, gridInfo GridInfo, node Node) map[string]Node {
	toCheck := nextToCheck(node, gridInfo)

	for len(toCheck) > 0 {
		var afterThat []Coordinate

		for _, v := range toCheck {
			tmp := gridData[v.toString()]
			if tmp.visited == true && gridInfo.startNodeAt.toString() != tmp.coordinate.toString() {
				continue
			}

			if tmp.value <= node.value+1 {
				fmt.Println("Comparing values: ", tmp.value, node.value)
				tmp.distanceFromStart = node.distanceFromStart + 1
				tmp.visited = true
				delete(gridData, v.toString())
				gridData[v.toString()] = tmp

				if tmp.coordinate.toString() != gridInfo.endNodeAt.toString() {
					afterThat = append(afterThat, nextToCheck(tmp, gridInfo)...)
					fmt.Println(afterThat, " are neighbours of: ", tmp)
				}
				node = tmp
			}
		}

		var filtered []Coordinate
		for _, v := range afterThat {
			if gridData[v.toString()].visited == false {
				filtered = append(filtered, v)
			}
		}

		toCheck = filtered
		fmt.Println("Neighbours to check: ", toCheck)
	}

	return gridData
}

func nextToCheck(node Node, gridInfo GridInfo) []Coordinate {
	var toCheck []Coordinate

	up, upE := nodeCoordinate(node, "up", gridInfo)
	if upE != "Error" {
		toCheck = append(toCheck, up)
	}

	down, downE := nodeCoordinate(node, "down", gridInfo)
	if downE != "Error" {
		toCheck = append(toCheck, down)
	}

	left, leftE := nodeCoordinate(node, "left", gridInfo)
	if leftE != "Error" {
		toCheck = append(toCheck, left)
	}

	right, rightE := nodeCoordinate(node, "right", gridInfo)
	if rightE != "Error" {
		toCheck = append(toCheck, right)
	}

	return toCheck
}

func nodeCoordinate(node Node, direction string, gridInfo GridInfo) (Coordinate, string) {
	switch direction {
	case "up":
		if node.coordinate.y+1 < gridInfo.sizeY {
			return Coordinate{node.coordinate.x, node.coordinate.y + 1}, ""
		}
	case "down":
		if node.coordinate.y-1 >= 0 {
			return Coordinate{node.coordinate.x, node.coordinate.y - 1}, ""
		}
	case "left":
		if node.coordinate.x+1 < gridInfo.sizeX {
			return Coordinate{node.coordinate.x + 1, node.coordinate.y}, ""
		}
	case "right":
		if node.coordinate.x-1 >= 0 {
			return Coordinate{node.coordinate.x - 1, node.coordinate.y}, ""
		}
	}
	return Coordinate{-1, -1}, "Error"
}

func readInputData(fileName string, gridInfo GridInfo) map[string]Node {
	inputData := make(map[string]Node, gridInfo.sizeX*gridInfo.sizeY)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for x, v := range line {
			coordinate := Coordinate{x, y}

			if string(v) == "S" {
				v = 97
			} else if string(v) == "E" {
				v = 122
			}
			node := Node{coordinate: coordinate, value: int(v)}

			inputData[coordinate.toString()] = node
		}
		y++
	}

	return inputData
}

func (receiver Coordinate) toString() string {
	return fmt.Sprintf("[%v,%v]", receiver.x, receiver.y)
}

func getGridInfo(fileName string) GridInfo {
	sizeX := 0
	sizeY := 0
	startCoordinate := Coordinate{}
	endCoordinate := Coordinate{}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for x, v := range line {
			coordinate := Coordinate{x, y}

			if string(v) == "E" {
				endCoordinate = coordinate
			} else if string(v) == "S" {
				startCoordinate = coordinate
			}
		}
		sizeX = len(line)
		y++
	}
	sizeY = y

	return GridInfo{sizeX: sizeX, sizeY: sizeY, startNodeAt: startCoordinate, endNodeAt: endCoordinate}
}
