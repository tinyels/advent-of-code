package lib

import (
	"testing"
)

func TestIsPossiblePassword(t *testing.T) {
	tables := []struct {
		candidate int
		output    bool
		output2   bool
	}{
		{111111, true, false},
		{223450, false, false},
		{123789, false, false},
		{112233, true, true},
		{123444, true, false},
		{111122, true, true},
	}

	for _, table := range tables {
		result, result2 := IsPossiblePassword(table.candidate)
		if result != table.output {
			t.Errorf("result was incorrect for %d, got %v, want %v", table.candidate, result, table.output)
		}
		if result2 != table.output2 {
			t.Errorf("result2 was incorrect for %d, got %v, want %v", table.candidate, result2, table.output2)
		}

	}
}
