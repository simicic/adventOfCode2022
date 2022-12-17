package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Coordinate struct {
	x int
	y int
}

type Node struct {
	coordinate Coordinate
	value      int
	visited    bool
}

type GridInfo struct {
	startNodeAt Coordinate
	endNodeAt   Coordinate
	sizeX       int
	sizeY       int
}

func main() {
	shortestPath := ShortestPath("12/input.txt")
	fmt.Println("Shortest path: ", shortestPath)

	shortestFromA := ShortestFromA("12/input.txt")
	fmt.Println("Shortest from any A: ", shortestFromA)
}

func ShortestFromA(fileName string) int {
	gridInfo := getGridInfo(fileName)
	gridData := readInputData(fileName, gridInfo)

	shortestPath := math.MaxInt
	for _, startNode := range gridData {
		if startNode.value == 'a' {
			cleanGridData(gridData)
			distance := visitPaths(startNode, gridData, gridInfo)

			if distance < shortestPath {
				shortestPath = distance
			}
		}
	}

	return shortestPath
}

func cleanGridData(gridData map[string]*Node) {
	for _, v := range gridData {
		v.visited = false
	}
}

func ShortestPath(fileName string) int {
	gridInfo := getGridInfo(fileName)
	gridData := readInputData(fileName, gridInfo)

	startNode := gridData[gridInfo.startNodeAt.toString()]
	shortestPath := visitPaths(startNode, gridData, gridInfo)

	return shortestPath
}

func visitPaths(currentNode *Node, gridData map[string]*Node, info GridInfo) int {
	pathsInLevels := make([][]*Node, 1000)
	pathsInLevels[0] = []*Node{currentNode}

	toVisit := nextToVisit(currentNode, gridData)
	currentNode.visited = true
	level := 1
	found := false

visiting:
	for len(toVisit) > 0 {
		nextVisit := make([]*Node, 0)

		for _, v := range toVisit {
			if v.visited {
				continue
			}
			v.visited = true
			nextVisit = append(nextVisit, nextToVisit(v, gridData)...)

			if v.coordinate.toString() == info.endNodeAt.toString() {
				found = true
				break visiting
			}
		}
		pathsInLevels[level] = toVisit
		toVisit = nextVisit
		level++
	}

	if found {
		return level
	} else {
		return math.MaxInt
	}
}

func nextToVisit(startNode *Node, gridData map[string]*Node) []*Node {
	var toVisit []*Node

	// left
	left := Coordinate{x: startNode.coordinate.x - 1, y: startNode.coordinate.y}
	leftElem, found := gridData[left.toString()]
	if found == true && leftElem.value-1 <= startNode.value && leftElem.visited == false {
		toVisit = append(toVisit, leftElem)
	}

	// right
	right := Coordinate{x: startNode.coordinate.x + 1, y: startNode.coordinate.y}
	rightElem, found := gridData[right.toString()]
	if found == true && rightElem.value-1 <= startNode.value && rightElem.visited == false {
		toVisit = append(toVisit, rightElem)
	}

	// up
	up := Coordinate{x: startNode.coordinate.x, y: startNode.coordinate.y + 1}
	upElem, found := gridData[up.toString()]
	if found == true && upElem.value-1 <= startNode.value && upElem.visited == false {
		toVisit = append(toVisit, upElem)
	}

	// down
	down := Coordinate{x: startNode.coordinate.x, y: startNode.coordinate.y - 1}
	downElem, found := gridData[down.toString()]
	if found == true && downElem.value-1 <= startNode.value && downElem.visited == false {
		toVisit = append(toVisit, downElem)
	}

	return toVisit
}

func readInputData(fileName string, gridInfo GridInfo) map[string]*Node {
	inputData := make(map[string]*Node, gridInfo.sizeX*gridInfo.sizeY)

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

			inputData[coordinate.toString()] = &node
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

			if v == 'E' {
				endCoordinate = coordinate
			} else if v == 'S' {
				startCoordinate = coordinate
			}
		}
		sizeX = len(line)
		y++
	}
	sizeY = y

	return GridInfo{sizeX: sizeX, sizeY: sizeY, startNodeAt: startCoordinate, endNodeAt: endCoordinate}
}
