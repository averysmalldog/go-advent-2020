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
	for k, v := range list {
		if v[0:3] == "nop"{
			//fmt.Printf("Test change on %d: ", k)
			var experiementalProgram []string
			_ = copy(experiementalProgram, list)
			// fmt.Printf("\t%s\n", experiementalProgram[k])
			experiementalProgram[k] = strings.ReplaceAll(experiementalProgram[k],"nop","jmp")
			// fmt.Printf("\t%s\n", experiementalProgram[k])
			accumulator := 0
			stepCounter := 1
			programCounter := 0
			instructionMemory := map[int]int{}
			badInstruction := compute(experiementalProgram, &instructionMemory, &programCounter, &accumulator, &stepCounter)
			//fmt.Printf("InstructionCache: %+v\n", instructionMemory)
			if badInstruction == -1 {
				fmt.Printf("Success! Instruction %d changed from nop to jmp, PC: %d, Accumulator: %d, Steps taken: %d\n", k, programCounter, accumulator, stepCounter)
				// fmt.Printf("InstructionCache: %+v\n", instructionMemory)
			} else {
				//fmt.Printf("Failure. Loops back to instruction %d.\n", badInstruction)
			}
		} else if v[0:3] == "jmp" {
			//fmt.Printf("Test change on %d: ", k)
			// fmt.Printf("Changing instruction %d from jmp to nop\n", k)
			experiementalProgram := list
			// fmt.Printf("\t%s\n", experiementalProgram[k])
			experiementalProgram[k] = strings.ReplaceAll(experiementalProgram[k],"jmp","nop")
			// fmt.Printf("\t%s\n", experiementalProgram[k])
			accumulator := 0
			stepCounter := 1
			programCounter := 0
			instructionMemory := map[int]int{}
			badInstruction := compute(experiementalProgram, &instructionMemory, &programCounter, &accumulator, &stepCounter)
			// fmt.Printf("InstructionCache: %+v\n", instructionMemory)
			if badInstruction == -1 {
				fmt.Printf("Success! Instruction %d changed from jmp to nop, PC: %d, Accumulator: %d, Steps taken: %d\n", k, programCounter, accumulator, stepCounter)
				// fmt.Printf("InstructionCache: %+v\n", instructionMemory)
			} else {
				//fmt.Printf("Failure. Loops back to instruction %d.\n", badInstruction)
			}
		} 
	}
}

func compute(program []string, instructionMemory *map[int]int, programCounter *int, accumulator *int, stepCounter *int) int {

	// if we've been here before, this is the loop
	localCache := *instructionMemory
	if localCache[*programCounter] != 0 {
		return *programCounter
	}

	// vars
	badInstruction := 0
	localCache[*programCounter] = *stepCounter
	*instructionMemory = localCache
	*stepCounter++

	// if we've reached the end of the program, exit
	if *programCounter == len(program)-1{
		return -1
	}

	// form up data
	instruction := strings.Fields(program[*programCounter])

	// case nop
	if instruction[0] == "nop" {
		// fmt.Printf("PC: %d, Instruction: %s, Accumulator: %d, Next op: ", *programCounter, instruction[0], *accumulator)
		*programCounter++
		// fmt.Println(*programCounter)
		return compute(program, instructionMemory, programCounter, accumulator, stepCounter)
	} else if instruction[0] == "acc" {
		// fmt.Printf("PC: %d, Instruction: %s, Accumulator: %d, Next op: ", *programCounter, instruction[0], *accumulator)
		if string(instruction[1][0]) == string("+") {
			value, _ := strconv.Atoi(string(instruction[1][1:]))
			*accumulator += value
			*programCounter++
			// fmt.Println(*programCounter)
			return compute(program, instructionMemory, programCounter, accumulator, stepCounter)
		}
		if string(instruction[1][0]) == string("-") {
			value, _ := strconv.Atoi(string(instruction[1][1:]))
			*accumulator -= value
			*programCounter++
			// fmt.Println(*programCounter)
			return compute(program, instructionMemory, programCounter, accumulator, stepCounter)
		}
	} else if instruction[0] == "jmp" {
		// fmt.Printf("PC: %d, Instruction: %s, Accumulator: %d, Next op: ", *programCounter, instruction[0], *accumulator)
		if string(instruction[1][0]) == string("+") {
			value, _ := strconv.Atoi(string(instruction[1][1:]))
			*programCounter += value
			// fmt.Println(*programCounter)
			return compute(program, instructionMemory, programCounter, accumulator, stepCounter)
		}
		if string(instruction[1][0]) == string("-") {
			value, _ := strconv.Atoi(string(instruction[1][1:]))
			*programCounter -= value
			// fmt.Println(*programCounter)
			return compute(program, instructionMemory, programCounter, accumulator, stepCounter)
		}
	}

	return badInstruction
}
