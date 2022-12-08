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
	totalSize := TotalSize("07/input.txt")
	fmt.Println("Total size: ", totalSize)
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

	calculateSizes(&rootDirectory)

	allSizes := storeAllSizes(&rootDirectory)

	minDelete := 30000000
	totalUnused := 70000000 - rootDirectory.size

	for _, v := range allSizes {
		if v < 100000 {
			totalSize += v
		}

		if v >= 30000000-totalUnused && v < totalUnused {
			if minDelete > v {
				minDelete = v
			}
		}
	}

	fmt.Println("Min size to delete: ", minDelete)
	return totalSize
}

func storeAllSizes(directory *Directory) []int {
	var allSizes []int

	allSizes = append(allSizes, directory.size)

	for _, dir := range directory.children {
		allSizes = append(allSizes, storeAllSizes(dir)...)
	}

	return allSizes
}

func calculateSizes(directory *Directory) int {
	filesSize := 0

	if len(directory.children) > 0 {
		for _, dir := range directory.children {
			filesSize += calculateSizes(dir)
		}
	}

	for _, v := range directory.files {
		filesSize += v.size
	}

	directory.size = filesSize

	return filesSize
}
