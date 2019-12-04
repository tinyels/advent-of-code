package lib

import (
	"testing"
)

func TestIsPossiblePassword(t *testing.T) {
	tables := []struct {
		candidate int
		output    bool
	}{
		{111111, true},
		{223450, false},
		{123789, false},
	}

	for _, table := range tables {
		result := IsPossiblePassword(table.candidate)
		if result != table.output {
			t.Errorf("result was incorrect for %d, got %v, want %v", table.candidate, result, table.output)
		}
	}
}
