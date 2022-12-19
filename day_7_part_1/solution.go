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
	Name string
	Size int
}

type Directory struct {
	Files    []File
	Children map[string]*Directory
	Parent   *Directory
}

func addDirectoryAndChildrenBelow100000(node *Directory, sum *int) int {
	size := 0
	for _, f := range (*node).Files {
		size += f.Size
	}
	for _, d := range (*node).Children {
		size += addDirectoryAndChildrenBelow100000(d, sum)
	}
	if size <= 100000 {
		*sum += size
	}
	return size
}

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var currentNode *Directory
	root := Directory{
		Files:    make([]File, 0),
		Children: make(map[string]*Directory),
	}
	currentNode = &root

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$") {
			command_parts := strings.Split(line, " ")
			if command_parts[1] == "cd" {
				if command_parts[2] == "/" {
					currentNode = &root
				} else if command_parts[2] == ".." {
					currentNode = (*currentNode).Parent
				} else {
					if _, ok := (*currentNode).Children[command_parts[2]]; ok {
						currentNode = (*currentNode).Children[command_parts[2]]
					} else {
						(*currentNode).Children[command_parts[2]] = &Directory{
							Files:    make([]File, 0),
							Children: make(map[string]*Directory),
							Parent:   currentNode,
						}
						currentNode = currentNode.Children[command_parts[2]]
					}
				}
			}
		} else if strings.HasPrefix(line, "dir") {
			dir_parts := strings.Split(line, " ")
			if _, ok := (*currentNode).Children[dir_parts[1]]; !ok {
				(*currentNode).Children[dir_parts[1]] = &Directory{
					Files:    make([]File, 0),
					Children: make(map[string]*Directory),
					Parent:   currentNode,
				}
			}
		} else {
			file_parts := strings.Split(line, " ")
			size, _ := strconv.Atoi(file_parts[0])
			currentNode.Files = append((*currentNode).Files, File{
				Name: file_parts[1],
				Size: size,
			})
		}
	}

	result := 0
	addDirectoryAndChildrenBelow100000(&root, &result)

	fmt.Println("Result:", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
