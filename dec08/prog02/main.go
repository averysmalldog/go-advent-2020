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
			// fmt.Printf("Changing instruction %d from nop to jmp\n", k)
			experiementalProgram := list
			// fmt.Printf("\t%s\n", experiementalProgram[k])
			experiementalProgram[k] = strings.ReplaceAll(experiementalProgram[k],"nop","jmp")
			// fmt.Printf("\t%s\n", experiementalProgram[k])
			accumulator := 0
			programCounter := 0
			instructionMemory := map[int]int{}
			badInstruction := compute(experiementalProgram, &instructionMemory, &programCounter, &accumulator)
			fmt.Printf("InstructionCache: %+v\n", instructionMemory)
			if badInstruction == -1 {
				fmt.Printf(" YES - Program terminated! Instruction %d changed from nop to jmp, PC: %d, Accumulator: %d\n", k, programCounter, accumulator)
			}
		} else if v[0:3] == "jmp" {
			// fmt.Printf("Changing instruction %d from jmp to nop\n", k)
			experiementalProgram := list
			// fmt.Printf("\t%s\n", experiementalProgram[k])
			experiementalProgram[k] = strings.ReplaceAll(experiementalProgram[k],"jmp","nop")
			// fmt.Printf("\t%s\n", experiementalProgram[k])
			accumulator := 0
			programCounter := 0
			instructionMemory := map[int]int{}
			badInstruction := compute(experiementalProgram, &instructionMemory, &programCounter, &accumulator)
			//fmt.Printf("InstructionCache: %+v\n", instructionMemory)
			if badInstruction == -1 {
				fmt.Printf(" YES - Program terminated! Instruction %d changed from jmp to nop, PC: %d, Accumulator: %d\n", k, programCounter, accumulator)
			} else {
				fmt.Printf(" badInstruction: %d", badInstruction)
			}
		}
	}
}

func compute(program []string, instructionMemory *map[int]int, programCounter *int, accumulator *int) int {
	// if we've reached the end of the program, exit
	if *programCounter == len(program){
		return -1
	}
	
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
		// fmt.Printf("PC: %d, Instruction: %s, Accumulator: %d, Next op: ", *programCounter, instruction[0], *accumulator)
		*programCounter++
		// fmt.Println(*programCounter)
		return compute(program, instructionMemory, programCounter, accumulator)
	}

	// case acc
	if instruction[0] == "acc" {
		// fmt.Printf("PC: %d, Instruction: %s, Accumulator: %d, Next op: ", *programCounter, instruction[0], *accumulator)
		if string(instruction[1][0]) == string("+") {
			value, _ := strconv.Atoi(string(instruction[1][1:]))
			*accumulator += value
			*programCounter++
			// fmt.Println(*programCounter)
			return compute(program, instructionMemory, programCounter, accumulator)
		}
		if string(instruction[1][0]) == string("-") {
			value, _ := strconv.Atoi(string(instruction[1][1:]))
			*accumulator -= value
			*programCounter++
			// fmt.Println(*programCounter)
			return compute(program, instructionMemory, programCounter, accumulator)
		}
	}

	// case jmp
	if instruction[0] == "jmp" {
		// fmt.Printf("PC: %d, Instruction: %s, Accumulator: %d, Next op: ", *programCounter, instruction[0], *accumulator)
		if string(instruction[1][0]) == string("+") {
			value, _ := strconv.Atoi(string(instruction[1][1:]))
			*programCounter += value
			// fmt.Println(*programCounter)
			return compute(program, instructionMemory, programCounter, accumulator)
		}
		if string(instruction[1][0]) == string("-") {
			value, _ := strconv.Atoi(string(instruction[1][1:]))
			*programCounter -= value
			// fmt.Println(*programCounter)
			return compute(program, instructionMemory, programCounter, accumulator)
		}
	}

	return badInstruction
}
