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
	start             bool
	end               bool
	distanceFromStart int
}

func main() {
	shortestPath := ShortestPath("12/input.txt")
	fmt.Println("Shortest path: ", shortestPath)
}

func ShortestPath(fileName string) int {
	gridData := readInputData(fileName)
	fmt.Println(gridData)

	return 0
}

func readInputData(fileName string) map[string]Node {
	inputData := make(map[string]Node)

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
			node := Node{coordinate: coordinate}

			if string(v) == "E" {
				node.end = true
				node.value = int(122)
			} else if string(v) == "S" {
				node.start = true
				node.value = int(97)
			} else {
				node.value = int(v)
			}

			inputData[fmt.Sprintf("[%v,%v]", x, y)] = node
		}
		x++
	}

	return inputData
}
