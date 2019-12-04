package main

import (
	"fmt"
	"./lib"
)

func main() {
	countForPartOne := 0
	for i :=240298; i <=784956; i++{
		if lib.IsPossiblePassword(i) {
			countForPartOne++
		}

	}
	fmt.Println("part 1: ", countForPartOne)
}
