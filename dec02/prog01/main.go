package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findAddsTo2020(list []int) (int, int, int) {
	for i1, n1 := range list {
		for i2, n2 := range list[i1+1:] {
			for i3, n3 := range list[i2+1:] {
				if n1+n2+n3 == 2020 {
					fmt.Printf("Found a match!\n\tIndex: %d, Value: %d\n\tIndex: %d, Value %d\n\tIndex: %d, Value %d\n", i1, n1, i2, n2, i3, n3)
					return n1, n2, n3
				}
			}
		}
	}
	return 0, 0, 0
}

func getSecretKey(int1 int, int2 int, int3 int) int {
	return int1 * int2 * int3
}

func main() {
	// test for presence of input.txt
	_, err := os.Stat("input.txt")
	if os.IsNotExist(err) {
		fmt.Println("input.txt not found. Exiting program.")
		return
	}

	// read input from the file!
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
	validCounter := 0

	for k, v := range list {
		fmt.Printf("Entry %d: %s: ", k, v)
		// Parse string into fields
		str := strings.Fields(v)

		minMax := strings.Split(str[0],"-")
		minOccurrences, _ := strconv.Atoi(minMax[0])
		maxOccurrences, _ := strconv.Atoi(minMax[1])
		desiredLetter := string(str[1][0])
		password := str[2]

		// Match for desired letter
		containsA := strings.Contains(password, desiredLetter)

		// get count if match
		if containsA {
			numInstances := strings.Count(password, desiredLetter)
			if (numInstances >= minOccurrences) && (numInstances <= maxOccurrences) {
				validCounter++
				fmt.Printf("valid!\n")
			} else {
				fmt.Printf("not valid\n")
			}
		}
	}

	fmt.Printf("Valid passwords: %d\n", validCounter)
	
}
