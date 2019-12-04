package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	
	"./lib"
)

func readFileAsIntArray(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var ints []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		if err != nil {
			log.Fatal(err)
		}
		for _, value := range values {
			value, err := strconv.Atoi((value))
			if err != nil {
				log.Fatal(err)
			}
			ints = append(ints, value)
		}
	}
	return ints, scanner.Err()
}

func partOne(input []int){
	program := make( []int, len(input))
	copy(program, input)
	program[1] = 12
	program[2] = 2
	output := lib.Intcode(program)
	fmt.Println("part 1:",output[0])
}
func partTwo(input []int){
	var noun, verb int
	for noun =0; noun <100 ; noun++ {
		for verb = 0; verb < 100; verb++ {
			program := make( []int, len(input))
			copy(program, input)
			program[1] = noun
			program[2] = verb
			output := lib.Intcode(program)
			if output[0] == 19690720 {
				fmt.Println("part 2:", 100*noun + verb)
			}
		}
	}
}
func main() {
	ints, err := readFileAsIntArray("input")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ints)
	partOne(ints)
	partTwo(ints)
}
