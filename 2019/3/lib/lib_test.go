package lib

import (
	"reflect"
	"testing"
)

func TestManhattanDistance(t *testing.T) {
	tables := []struct {
		a      Coordinate
		b      Coordinate
		output int
	}{
		{Coordinate{0, 0}, Coordinate{0, 0}, 0},
		{Coordinate{0, 0}, Coordinate{3, 3}, 6},
		{Coordinate{3, 3}, Coordinate{0, 0}, 6},
	}

	for _, table := range tables {
		result := ManhattanDistance(table.a, table.b)
		if result != table.output {
			t.Errorf("result was incorrect, got %d, want %d", result, table.output)
		}
	}
}

func TestGetWireCoordinates(t *testing.T) {
	tables := []struct {
		a            Coordinate
		instructions string
		output       []Coordinate
	}{
		{Coordinate{0, 0}, "U1", []Coordinate{Coordinate{0, 1}}},
		{Coordinate{0, 0}, "R1", []Coordinate{Coordinate{1, 0}}},
		{Coordinate{1, 0}, "L1", []Coordinate{Coordinate{0, 0}}},
		{Coordinate{1, 1}, "D1", []Coordinate{Coordinate{1, 0}}},
		{Coordinate{5, 10}, "D4", []Coordinate{Coordinate{5, 9}, Coordinate{5, 8}, Coordinate{5, 7}, Coordinate{5, 6}}},
		{Coordinate{1, 1}, "D1,R1", []Coordinate{Coordinate{1, 0}, Coordinate{2, 0}}},
	}

	for _, table := range tables {
		result := GetWireCoordinates(table.a, table.instructions)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("result was incorrect, got %d, want %d", result, table.output)
		}
	}
}

func TestSolve(t *testing.T) {
	tables := []struct {
		a      string
		b      string
		output int
	}{
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}

	for _, table := range tables {
		result := SolvePartOne(table.a, table.b)
		if result != table.output {
			t.Errorf("result was incorrect, got %d, want %d", result, table.output)
		}
	}
}
func TestGetFewestSteps(t *testing.T) {
	tables := []struct {
		a      string
		b      string
		output int
	}{
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 610},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 410},
	}

	for _, table := range tables {
		result := GetFewestSteps(table.a, table.b)
		if result != table.output {
			t.Errorf("result was incorrect, got %d, want %d", result, table.output)
		}
	}
}
