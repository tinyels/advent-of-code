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

func main() {
	ints, err := readFileAsIntArray("input")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ints)
	ints[1] = 12
	ints[2] = 2
	output := lib.Intcode(ints)
	fmt.Println(output)
}
