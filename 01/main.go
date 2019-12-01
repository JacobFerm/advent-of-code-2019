package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := readInput()

	baseFuel := calcBaseFuel(input)
	fmt.Printf("Initial fuel usage: %d\n", baseFuel)

	totalFuel := calcTotalFuel(input)
	fmt.Printf("Total fuel usage: %d\n", totalFuel)
}

func readInput() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []int
	for scanner.Scan() {
		input = append(input, readMass(scanner.Text()))
	}
	return input
}

func readMass(input string) int {
	mass, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return mass
}

func calcBaseFuel(input []int) int {
	fuel := 0
	for _, v := range input {
		fuel += calcFuel(v)
	}
	return fuel
}

func calcFuel(mass int) int {
	return mass/3 - 2
}

func calcTotalFuel(input []int) int {
	fuel := 0
	for _, v := range input {
		fuel += calcFuelForModule(v)
	}
	return fuel
}

func calcFuelForModule(mass int) int {
	sum := 0
	newFuel := calcFuel(mass)
	for newFuel > 0 {
		sum += newFuel
		newFuel = calcFuel(newFuel)
	}
	return sum
}
