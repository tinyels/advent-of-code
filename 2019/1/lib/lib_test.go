package lib

import "testing"

func TestFuelForLaunch(t *testing.T) {
	tables := []struct {
		mass int
		fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, table := range tables {
		result := FuelRequiredToLaunch(table.mass)
		if result != table.fuel {
			t.Errorf("result was incorrect, got %d, want %d", result, table.fuel)
		}
	}
}
func TestInclusiveFuelForLaunch(t *testing.T) {
	tables := []struct {
		mass int
		fuel int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, table := range tables {
		result := InclusiveFuelRequiredToLaunch(table.mass)
		if result != table.fuel {
			t.Errorf("result was incorrect, got %d, want %d", result, table.fuel)
		}
	}
}
