package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	fmt.Printf("Bad Instruction: %d, Accumulator: %d\n", badInstruction, accumulator)
}

func compute(instructionList []string, instructionMemory *map[int]int, programCounter *int, accumulator *int) int {
	badInstruction := *programCounter
	if *programCounter < 10 {
		fmt.Printf("Hello World. %d\n", *programCounter)
		*programCounter++
		badInstruction = compute(instructionList, instructionMemory, programCounter, accumulator)
	}
	
	return badInstruction
}