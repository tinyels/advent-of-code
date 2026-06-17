package main

import (
	"fmt"
	"./lib"
)

func main() {
	countForPartOne := 0
	countForPartTwo := 0
	for i :=240298; i <=784956; i++{
		a, b := lib.IsPossiblePassword(i) 
		if a {
			countForPartOne++
		}
		if b {
			countForPartTwo++
		}

	}
	fmt.Println("part 1: ", countForPartOne) //1150
	fmt.Println("part 2: ", countForPartTwo) //748
}
