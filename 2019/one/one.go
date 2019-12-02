package one

func FuelRequiredToLaunch(mass int) int {
	f := (mass / 3) - 2
	if f < 0 {
		return 0
	}
	return f
}

func InclusiveFuelRequiredToLaunch(mass int) int {
	if mass <= 0 {
		return 0
	}
	f := FuelRequiredToLaunch(mass)
	return f + InclusiveFuelRequiredToLaunch(f)

}
