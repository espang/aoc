package day14

import (
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func Add(p Point, delta Point) Point {
	p.x += delta.x
	p.y += delta.y
	return p
}

func p(x, y int) Point {
	return Point{x: x, y: y}
}

func parsePoint(point string) Point {
	splitted := strings.Split(point, ",")
	if len(splitted) != 2 {
		panic("parsePoint failed: " + point)
	}
	x, err1 := strconv.Atoi(splitted[0])
	y, err2 := strconv.Atoi(splitted[1])
	if err1 != nil || err2 != nil {
		panic("parsePoint failed: " + point)
	}
	return p(x, y)
}

func parseLine(line string) []Point {
	rawPoints := strings.Split(line, " -> ")
	points := []Point{}
	for _, rawPoint := range rawPoints {
		points = append(points, parsePoint(rawPoint))
	}
	return points
}

func parseInput(s string) [][]Point {
	lines := strings.Split(s, "\n")
	ppoints := [][]Point{}
	for _, line := range lines {
		ppoints = append(ppoints, parseLine(line))
	}
	return ppoints
}
