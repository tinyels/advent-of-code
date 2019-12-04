package main

import (
	"../util"
	"./lib"
	"bufio"
	"fmt"
	"os"
)

func readFile(path string) []string {
	file, err := os.Open(path)
	util.Check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func partOne(input []string) {
	fmt.Println("part 1: ", lib.GetClosestDistance(input[0], input[1]))
}
func partTwo(input []string) {
	fmt.Println("part 2:", lib.GetFewestSteps(input[0], input[1]))
}
func main() {
	lines := readFile("input")
	partOne(lines)
	partTwo(lines)
}
