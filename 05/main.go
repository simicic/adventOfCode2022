package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	topCratesCode := TopCratesCode("05/input.txt")
	fmt.Println("Top crates code: ", topCratesCode)
}

func TopCratesCode(fileName string) string {
	code := ""

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	return code
}
