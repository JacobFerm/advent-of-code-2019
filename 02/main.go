package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := readInput()

	fmt.Printf("1202: %d\n", intcode(input, 12, 2))

	param1 := 12
	param2 := 2

	for intcode(input, param1, param2) < 19690720 {
		param1 += 1
	}
	param1 -= 1
	for intcode(input, param1, param2) <= 19690720 {
		param2 += 1
	}
	param2 -= 1

	fmt.Printf("param1:%d  param2:%d = %d\n", param1, param2, intcode(input, param1, param2))
}

func readInput() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	record, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	input := make([]int, len(record))

	for i, v := range record {
		op, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		input[i] = op
	}
	return input
}

func intcode(input []int, param1 int, param2 int) int {
	memory := make([]int, len(input))
	copy(memory, input)
	memory[1] = param1
	memory[2] = param2
	for i := 0; memory[i] != 99; i += 4 {
		switch memory[i] {
		case 1:
			memory[memory[i+3]] = memory[memory[i+1]] + memory[memory[i+2]]
		case 2:
			memory[memory[i+3]] = memory[memory[i+1]] * memory[memory[i+2]]
		default:
			fmt.Printf("Invalid opcode %d on position %d\n", memory[i], i)
		}
	}
	return memory[0]
}
