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
	coordinate  Coordinate
	passThrough bool
}

//type GridInfo struct {
//	startNodeAt Coordinate
//	endNodeAt   Coordinate
//	sizeX       int
//	sizeY       int
//}

func main() {
	finalPassword := FinalPassword("22/input_test.txt")
	fmt.Println("Final password: ", finalPassword)
}

func FinalPassword(fileName string) int {
	data, commands := readInputData(fileName)

	fmt.Println(data)
	fmt.Println(commands)

	return 0
}

func readInputData(fileName string) ([]*Node, string) {
	var inputData []*Node
	var coordinates []*Coordinate
	commands := ""
	maxX := 0
	maxY := 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	y := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for x, v := range line {
			if len(line) == 0 {
				scanner.Scan()
				commands = scanner.Text()
				break
			}

			if v == ' ' {
				continue
			}

			coordinate := Coordinate{x + 1, y}
			coordinates = append(coordinates, &coordinate)

			passTrough := false
			if v == '.' {
				passTrough = true
			}
			node := Node{coordinate: coordinate, passThrough: passTrough}
			inputData = append(inputData, &node)
		}
		y++
	}

	return inputData, commands
}

func (receiver Coordinate) toString() string {
	return fmt.Sprintf("[%v,%v]", receiver.x, receiver.y)
}
