package lib

import "strconv"
import "strings"

type Coordinate struct {
	x int
	y int
}

func SolvePartOne(a string, b string) int {
	origin := Coordinate{0, 0}
	aWire := GetWireCoordinates(origin, a)
	bWire := GetWireCoordinates(origin, b)
	return GetClosestIntersectionDistance(aWire, bWire)
}

func ManhattanDistance(a Coordinate, b Coordinate) int {
	return abs(a.y-b.y) + abs(a.x-b.x)
}

func GetWireCoordinates(startingCoordinates Coordinate, instructions string) []Coordinate {
	var wire []Coordinate
	var pos Coordinate
	pos = startingCoordinates

	for _, instruction := range strings.Split(instructions, ",") {
		direction := instruction[:1]
		distance := getInt(instruction[1:])

		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				pos.y++
			case "D":
				pos.y--
			case "R":
				pos.x++
			case "L":
				pos.x--
			default:
				panic("unknown direction")
			}
			wire = append(wire, pos)
		}
	}
	return wire
}

func GetClosestIntersectionDistance(a []Coordinate, b []Coordinate) int {
	origin := Coordinate{0, 0}
	shortestDistance := -1
	for _, item := range Intersection(a, b) {
		if shortestDistance == -1 {
			shortestDistance = ManhattanDistance(origin, item)
		} else {
			shortestDistance = min(shortestDistance, ManhattanDistance(origin, item))
		}
	}
	return shortestDistance
}

func GetFewestSteps(a string, b string) int {
	origin := Coordinate{0, 0}
	aWire := GetWireCoordinates(origin, a)
	bWire := GetWireCoordinates(origin, b)
	steps := -1
	for _, item := range Intersection(aWire, bWire) {
		if steps == -1 {
			steps = 2 + indexOf(item, aWire) + indexOf(item, bWire)
		} else {
			steps = min(steps, 2+indexOf(item, aWire)+indexOf(item, bWire))
		}
	}
	return steps

}
func Intersection(a []Coordinate, b []Coordinate) []Coordinate {
	set := make([]Coordinate, 0)
	for _, item := range a {
		if contains(item, b) {
			set = append(set, item)
		}
	}
	return set
}

func contains(c Coordinate, list []Coordinate) bool {
	for _, a := range list {
		if a == c {
			return true
		}
	}
	return false
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getInt(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func indexOf(element Coordinate, data []Coordinate) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
