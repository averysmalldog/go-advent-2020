package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// test for presence of input.txt
	_, err := os.Stat("input.txt")
	if os.IsNotExist(err) {
		fmt.Println("input.txt not found. Exiting program.")
		return
	}

	// read input from the file and store as an array of strings
	path := "input.txt"
	var list []string

	buf, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	lineScanner := bufio.NewScanner(buf)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		if err != nil {
			log.Println("well I guess there's an error")
		}
		list = append(list, line)
	}
	err = lineScanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	// solution
	accumulator := 0
	programCounter := 0
	instructionMemory := map[int]int{}
	badInstruction := compute(list, &instructionMemory, &programCounter, &accumulator)
	fmt.Printf("InstructionCache: %+v\n", instructionMemory)
	fmt.Printf("Bad Instruction: %d, Accumulator: %d\n", badInstruction, accumulator)
}

func compute(program []string, instructionMemory *map[int]int, programCounter *int, accumulator *int) int {
	// if we've been here before, this is the loop
	localCache := *instructionMemory
	if localCache[*programCounter] != 0 {
		return *programCounter
	}

	// vars
	badInstruction := 0
	localCache[*programCounter]++
	*instructionMemory = localCache

	// form up data
	instruction := strings.Fields(program[*programCounter])

	// case nop
	if instruction[0] == "nop" {
		fmt.Printf("PC: %d, Instruction: %s, Accumulator: %d\n", *programCounter, instruction[0], *accumulator)
		*programCounter++
		return compute(program, instructionMemory, programCounter, accumulator)
	}

	// case acc
	if instruction[0] == "acc" {
		fmt.Printf("PC: %d, Instruction: %s, Accumulator: %d\n", *programCounter, instruction[0], *accumulator)
		if string(instruction[1][0]) == string("+") {
			value, err := strconv.Atoi(string(instruction[1][1:]))
			if err != nil {
				return -1
			}
			*accumulator += value
			*programCounter++
			return compute(program, instructionMemory, programCounter, accumulator)
		}
		if string(instruction[1][0]) == string("-") {
			value, err := strconv.Atoi(string(instruction[1][1:]))
			if err != nil {
				return -1
			}
			*accumulator -= value
			*programCounter++
			return compute(program, instructionMemory, programCounter, accumulator)
		}
	}

	// case jmp
	if instruction[0] == "jmp" {
		fmt.Printf("PC: %d, Instruction: %s, Accumulator: %d\n", *programCounter, instruction[0], *accumulator)
		if string(instruction[1][0]) == string("+") {
			value, err := strconv.Atoi(string(instruction[1][1:]))
			if err != nil {
				return -1
			}
			*programCounter += value
			return compute(program, instructionMemory, programCounter, accumulator)
		}
		if string(instruction[1][0]) == string("-") {
			value, err := strconv.Atoi(string(instruction[1][1:]))
			if err != nil {
				return -1
			}
			*programCounter -= value
			return compute(program, instructionMemory, programCounter, accumulator)
		}
	}

	return badInstruction
}
