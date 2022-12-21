package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Node struct {
	value    int
	next     *Node
	previous *Node
}

type List struct {
	head *Node
}

func main() {
	// 332 too low
	sum := SumGroveCoordinates("20/input.txt")
	fmt.Println("Sum of grove coordinates: ", sum)
}

func SumGroveCoordinates(fileName string) int {
	inputData, _ := readInputData(fileName)
	var zeroNode *Node
	printInputData(inputData)

	size := len(inputData)
	for _, v := range inputData {
		//fmt.Println("Looking to move: ", v.value)
		//
		nodeB := v
		nodeA := v

		if v.value == 0 {
			zeroNode = v
			continue
		} else if v.value > 0 {
			x := 0
			for x < v.value%size {
				nodeA = nodeA.next
				x++
			}
		} else if v.value < 0 {
			x := v.value
			for x <= 0 {
				nodeA = nodeA.previous
				x++
			}
		}

		nodeC := nodeA.next

		nodeAA := nodeB.previous
		nodeCC := nodeB.next

		nodeAA.next = nodeCC
		//fmt.Println("Node with value (AA): ", nodeAA.value, "has previous / next ", nodeAA.previous.value, nodeAA.next.value)

		nodeCC.previous = nodeAA
		//fmt.Println("Node with value (CC): ", nodeCC.value, "has previous / next ", nodeCC.previous.value, nodeCC.next.value)

		nodeA.next = nodeB
		//fmt.Println("Node with value (A): ", nodeA.value, "has previous / next ", nodeA.previous.value, nodeA.next.value)

		nodeC.previous = nodeB
		//fmt.Println("Node with value (C): ", nodeC.value, "has previous / next ", nodeC.previous.value, nodeC.next.value)

		nodeB.next = nodeC
		nodeB.previous = nodeA
		//fmt.Println("Node with value (B): ", nodeB.value, "has previous / next ", nodeB.previous.value, nodeB.next.value)

		//printList(&list)
		//fmt.Println()
	}

	// 1000th, 2000th, and 3000th

	sum := 0
	for i := 0; i < size; i++ {
		if 1000%size == i || 2000%size == i || 3000%size == i {
			sum += zeroNode.value
		}
		zeroNode = zeroNode.next
	}

	return sum
}

func printInputData(inputData []*Node) {
	for _, v := range inputData {
		fmt.Print("-> ", v.value)
	}
	fmt.Println()
}

func printList(l *List) {
	p := l.head
	for p != nil {
		fmt.Printf("-> %v ", p.value)
		p = p.next

		if p == l.head {
			break
		}
	}
}

func readInputData(fileName string) ([]*Node, List) {
	var inputData []*Node
	var list List

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var previous *Node
	var lastAdded *Node

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		lastAdded = list.Insert(value, previous)
		inputData = append(inputData, lastAdded)
		previous = lastAdded
	}

	lastAdded.next = list.head
	list.head.previous = lastAdded

	return inputData, list
}

func (l *List) Insert(d int, previous *Node) *Node {
	node := &Node{value: d, next: nil, previous: previous}
	p := l.head

	if l.head == nil {
		l.head = node
	} else {
		for p.next != nil {
			p = p.next
		}
		p.next = node
	}

	return node
}
