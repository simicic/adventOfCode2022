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

	fmt.Println(gridInfo)
	fmt.Println(gridData)

	return 0
}

func readInputData(fileName string, gridInfo GridInfo) [][]Node {
	inputData := make([][]Node, gridInfo.sizeX)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rowData := make([]Node, gridInfo.sizeY)

		for y, v := range line {
			coordinate := Coordinate{x, y}
			node := Node{coordinate: coordinate, value: int(v)}

			rowData[y] = node
		}
		inputData[x] = rowData
		x++
	}

	return inputData
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

	x := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for y, v := range line {
			coordinate := Coordinate{x, y}

			if string(v) == "E" {
				endCoordinate = coordinate
			} else if string(v) == "S" {
				startCoordinate = coordinate
			}
		}
		sizeY = len(line)
		x++
	}
	sizeX = x

	return GridInfo{sizeX: sizeX, sizeY: sizeY, startNodeAt: startCoordinate, endNodeAt: endCoordinate}
}
