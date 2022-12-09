package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Knot struct {
	x int
	y int
}

var head Knot
var tail Knot

func main() {
	visitedCount := VisitedCount("09/input.txt")
	fmt.Println("Visible knots: ", visitedCount)
}

func VisitedCount(fileName string) int {
	visitedCount := 0

	// move head (location)
	// follow with tail (location)
	// which places have been visited (slice of [x,y])

	head = Knot{0, 0}
	tail = Knot{0, 0}

	tailVisited := make(map[string]int)
	tailVisited["[0,0]"] = 1

	inputData := readInputData(fileName)

	for _, v := range inputData {
		direction, step := decipherCommand(v)

		move(&head, &tail, &tailVisited, direction, step)
	}
	visitedCount = len(tailVisited)

	return visitedCount
}

func move(head *Knot, tail *Knot, tailVisited *map[string]int, direction string, step int) {
	switch direction {
	case "R":
		for i := 1; i <= step; i++ {
			prevHead := *head
			head.x = head.x + 1

			actionTail(head, tail, prevHead, tailVisited)
		}
	case "L":
		for i := 1; i <= step; i++ {
			prevHead := *head
			head.x = head.x - 1

			actionTail(head, tail, prevHead, tailVisited)
		}
	case "U":
		for i := 1; i <= step; i++ {
			prevHead := *head
			head.y = head.y + 1

			actionTail(head, tail, prevHead, tailVisited)
		}
	case "D":
		for i := 1; i <= step; i++ {
			prevHead := *head
			head.y = head.y - 1

			actionTail(head, tail, prevHead, tailVisited)
		}
	}
}

func actionTail(head *Knot, tail *Knot, prevHead Knot, tailVisited *map[string]int) {
	if shouldMove(head, tail) {
		tail.moveTo(prevHead)
		coordinates := fmt.Sprintf("[%v,%v]", tail.x, tail.y)
		(*tailVisited)[coordinates]++
	}
}

func (c *Knot) moveTo(cell Knot) {
	c.x = cell.x
	c.y = cell.y
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

func decipherCommand(command string) (string, int) {
	slice := strings.Split(command, " ")
	step, _ := strconv.Atoi(slice[1])

	return slice[0], step
}

func shouldMove(head *Knot, tail *Knot) bool {
	distanceSquared := (head.x-tail.x)*(head.x-tail.x) + (head.y-tail.y)*(head.y-tail.y)

	if distanceSquared > 2 {
		return true
	} else {
		return false
	}
}
