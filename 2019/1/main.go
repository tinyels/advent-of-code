package main

import (
	"./lib"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readLinesAsInt(path string) ([]int, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var ints []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
        ints = append(ints, value)
    }
    return ints, scanner.Err()
}
func dayOne() {
	totalFuelRequired :=0
	totalInclusiveFuel := 0
	values, err := readLinesAsInt("one/input")
	if (err != nil) {
		log.Fatal(err)
	}
	for _, value := range values {
		totalFuelRequired += lib.FuelRequiredToLaunch(value)
		totalInclusiveFuel += lib.InclusiveFuelRequiredToLaunch(value)
	}
	fmt.Println(totalFuelRequired)
	fmt.Println(totalInclusiveFuel)

}
func main(){
	dayOne();
}
