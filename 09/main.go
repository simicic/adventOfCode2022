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
	x     int
	y     int
	steps []string
	value string
}

var head Knot
var tail Knot

func main() {
	visitedCount := VisitedCount("09/input.txt")
	fmt.Println("Visible knots: ", visitedCount)

	visitedCountLong := VisitedCountLong("09/input.txt")
	fmt.Println("Visible knots: ", visitedCountLong)
}

func VisitedCountLong(fileName string) int {
	visitedCount := 0

	inputData := readInputData(fileName)

	head := make([]*Knot, 10)

	for i := 0; i < 10; i++ {
		knot := Knot{x: 0, y: 0, value: fmt.Sprintf("%v", i)}
		head[i] = &knot
	}
	head[0].value = "H"

	for _, v := range inputData {
		direction, step := decipherCommand(v)
		moveLong(head, direction, step)
	}

	visitedCount = len(unique(head[9].steps)) + 1 // correct for the start element

	return visitedCount
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func moveTheRestOfTail(head []*Knot) {
	prevHead := *head[0]

	for i := 1; i < 10; i++ {
		if shouldMove(&prevHead, head[i]) {
			x, y := whereToMove(prevHead, head[i])
			head[i].x = x
			head[i].y = y
			head[i].steps = append(head[i].steps, fmt.Sprintf("[%v,%v]", x, y))
		}
		prevHead = *head[i]
	}
}

func whereToMove(prevHead Knot, knotToMove *Knot) (int, int) {
	deltaX := clamp(prevHead.x-knotToMove.x, -1, 1)
	deltaY := clamp(prevHead.y-knotToMove.y, -1, 1)

	return knotToMove.x + deltaX, knotToMove.y + deltaY
}

// https://en.wikipedia.org/wiki/Clamping_(graphics)
func clamp(v int, min int, max int) int {
	if v < min {
		return min
	} else if v > max {
		return max
	} else {
		return v
	}
}

func moveLong(head []*Knot, direction string, step int) {
	switch direction {
	case "R":
		for i := 0; i < step; i++ {
			head[0].x = head[0].x + 1

			moveTheRestOfTail(head)
		}
	case "L":
		for i := 0; i < step; i++ {
			head[0].x = head[0].x - 1

			moveTheRestOfTail(head)
		}
	case "U":
		for i := 0; i < step; i++ {
			head[0].y = head[0].y + 1

			moveTheRestOfTail(head)
		}
	case "D":
		for i := 0; i < step; i++ {
			head[0].y = head[0].y - 1

			moveTheRestOfTail(head)
		}
	}
}

func VisitedCount(fileName string) int {
	visitedCount := 0

	// move head (location)
	// follow with tail (location)
	// which places have been visited (slice of [x,y])

	head = Knot{x: 0, y: 0}
	tail = Knot{x: 0, y: 0}

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

func printOutcome(head []*Knot) {
	for j := 22; j >= -22; j-- {
		line := ""
		for i := -22; i <= 20; i++ {
			elem := "."
			for _, v := range head {
				if v.x == i && v.y == j {
					elem = v.value
					break
				}
			}
			line = line + elem
		}
		fmt.Println(line)
	}
}
