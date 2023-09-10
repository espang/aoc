package aoc

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type dir struct {
	parent *dir
	size   int
	name   string
	dirs   []*dir
	files  []file
}

type changeDirectory struct {
	to string
}

type list struct{}

type file struct {
	name string
	size int
}

func parseMessage(s string) any {
	if s == "$ ls" {
		return list{}
	}
	if strings.HasPrefix(s, "$ cd ") {
		return changeDirectory{to: strings.TrimPrefix(s, "$ cd ")}
	}
	splitted := strings.Split(s, " ")
	if len(splitted) != 2 {
		panic("unsupported message")
	}
	if splitted[0] == "dir" {
		return &dir{name: splitted[1]}
	}
	size, err := strconv.Atoi(splitted[0])
	if err != nil {
		panic("unsupported message: " + err.Error())
	}
	return file{name: splitted[1], size: size}
}

func fs(input string) *dir {
	lines := strings.Split(input, "\n")
	var root *dir
	var currentDir *dir

outer:
	for _, line := range lines {
		switch cmd := parseMessage(line).(type) {
		case changeDirectory:
			if cmd.to == "/" {
				root = &dir{name: "/"}
				currentDir = root
			} else if cmd.to == ".." {
				currentDir = currentDir.parent
			} else {
				for _, d := range currentDir.dirs {
					if d.name == cmd.to {
						currentDir = d
						continue outer
					}
				}
				panic("didn't find directory")
			}
		case *dir:
			cmd.parent = currentDir
			currentDir.dirs = append(currentDir.dirs, cmd)
		case file:
			currentDir.files = append(currentDir.files, cmd)
		case list:
			//skip
		}
	}
	return root
}

func calculateSizes(root *dir) {
	sizeOfFiles := 0
	for _, f := range root.files {
		sizeOfFiles += f.size
	}
	sizeOfDirectories := 0
	for _, child := range root.dirs {
		calculateSizes(child)
		sizeOfDirectories += child.size
	}
	root.size = sizeOfFiles + sizeOfDirectories
}

func sumOfAllBelow100k(d *dir) int {
	total := 0
	for _, child := range d.dirs {
		total += sumOfAllBelow100k(child)
	}
	if d.size <= 100_000 {
		total += d.size
	}
	return total
}

func Day7Part1(input string) {
	root := fs(input)
	calculateSizes(root)
	fmt.Println(sumOfAllBelow100k(root))
}

func allSizesBigEnough(d *dir, minimum int) []int {
	if d.size < minimum {
		return []int{}
	}
	sizes := []int{d.size}
	for _, child := range d.dirs {
		sizes = append(sizes, allSizesBigEnough(child, minimum)...)
	}
	return sizes
}

func Day7Part2(input string) {
	root := fs(input)
	calculateSizes(root)
	total := 70_000_000
	needed := 30_000_000
	toFree := needed - (total - root.size)
	sizes := allSizesBigEnough(root, toFree)
	sort.Ints(sizes)
	fmt.Println(sizes[0])
}
