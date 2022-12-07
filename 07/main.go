package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type Directory struct {
	name     string
	files    []*File
	children []*Directory
	parent   *Directory
	size     int
}

func main() {
	totalSize := TotalSize("07/input_test.txt")
	fmt.Println("Packet starts at index: ", totalSize)
}

func TotalSize(fileName string) int {
	totalSize := 0
	var rootDirectory Directory
	var currentDirectory *Directory

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "$ cd /") {
			rootDirectory = Directory{name: "/"}
			currentDirectory = &rootDirectory
			continue
		}

		if strings.HasPrefix(line, "$ cd") {
			if strings.HasPrefix(line, "$ cd ..") {
				currentDirectory = currentDirectory.parent
				continue
			} else {
				name := strings.Replace(line, "$ cd ", "", 1)

				for _, directory := range currentDirectory.children {
					if directory.name == name {
						currentDirectory = directory
						break
					}
				}
				continue
			}
		} else if strings.HasPrefix(line, "$ ls") {
			// nothing
			continue
		} else {
			file := strings.Split(line, " ")

			if file[0] == "dir" {
				newDirectory := Directory{name: strings.Replace(file[1], "dir ", "", 1), parent: currentDirectory}
				currentDirectory.children = append(currentDirectory.children, &newDirectory)
				continue
			} else {
				fileSize, _ := strconv.Atoi(file[0])
				newFile := File{name: file[1], size: fileSize}
				currentDirectory.files = append(currentDirectory.files, &newFile)
				continue
			}
		}
	}

	fmt.Println(rootDirectory)
	return totalSize
}
