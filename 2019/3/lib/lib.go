package lib

import "strings"
import "../../util"

type Coordinate struct {
	x int
	y int
}

func GetClosestDistance(a string, b string) int {
	origin := Coordinate{0, 0}
	aWire := GetWireCoordinates(origin, a)
	bWire := GetWireCoordinates(origin, b)
	return getClosestIntersectionDistance(aWire, bWire)
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
			steps = util.MinInt(steps, 2+indexOf(item, aWire)+indexOf(item, bWire))
		}
	}
	return steps
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
		distance := util.GetIntFromString(instruction[1:])

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

func getClosestIntersectionDistance(a []Coordinate, b []Coordinate) int {
	origin := Coordinate{0, 0}
	shortestDistance := -1
	for _, item := range Intersection(a, b) {
		if shortestDistance == -1 {
			shortestDistance = ManhattanDistance(origin, item)
		} else {
			shortestDistance = util.MinInt(shortestDistance, ManhattanDistance(origin, item))
		}
	}
	return shortestDistance
}

func Intersection(a []Coordinate, b []Coordinate) []Coordinate {
	set := make([]Coordinate, 0)
	hash := make(map[Coordinate]bool)
	for _, item := range a {
		hash[item] = true

	}
	for _, item := range b {
		if _, found := hash[item]; found {
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

func indexOf(element Coordinate, data []Coordinate) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
