package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	message := readFromFile("06/input.txt")
	fmt.Println("Packet starts at index: ", PacketStartIndex(message))
	fmt.Println("Message starts at index: ", MessageStartIndex(message))
}

func readFromFile(fileName string) string {
	message := ""

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		message = scanner.Text()
	}

	return message
}

func PacketStartIndex(message string) int {
	packetStart := 0

	for i := 0; i < len(message)-4; i++ {
		if message[i] == message[i+1] ||
			message[i] == message[i+2] ||
			message[i] == message[i+3] ||
			message[i+1] == message[i+2] ||
			message[i+1] == message[i+3] ||
			message[i+2] == message[i+3] {
			continue
		} else {
			packetStart = i + 4
			break
		}
	}
	return packetStart
}

func MessageStartIndex(message string) int {
	messageStart := 0
	uniquenessMap := make(map[rune]int, 14)

	for i := 0; i < len(message)-15; i++ {
		for j := i; j < i+14; j++ {
			uniquenessMap[rune(message[j])] += 1
		}

		unique := true
		for _, v := range uniquenessMap {
			if v != 1 {
				unique = false
			}
		}

		if unique == true {
			messageStart = i + 14
			break
		} else {
			uniquenessMap = make(map[rune]int, 14)
		}
	}
	return messageStart
}
