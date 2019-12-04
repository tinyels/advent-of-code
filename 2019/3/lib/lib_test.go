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
	}

	for _, table := range tables {
		result := GetWireCoordinates(table.a, table.instructions)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("result was incorrect, got %d, want %d", result, table.output)
		}
	}
}
