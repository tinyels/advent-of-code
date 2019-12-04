package lib

import "strconv"

type Coordinate struct {
	x int
	y int
}

func ManhattanDistance(a Coordinate, b Coordinate) int {
	return abs(a.y-b.y) + abs(a.x-b.x)
}

func GetWireCoordinates(startingCoordinates Coordinate, instruction string) []Coordinate {
	direction := instruction[:1]
	distance := getInt(instruction[1:])

	var wire []Coordinate
	var pos Coordinate
	pos = startingCoordinates

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
	return wire
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
