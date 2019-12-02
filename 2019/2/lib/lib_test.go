package lib

import (
	"reflect"
	"testing"
)

func TestFuelForLaunch(t *testing.T) {
	tables := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, table := range tables {
		result := Intcode(table.input)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("result was incorrect, got %d, want %d", result, table.output)
		}
	}
}
