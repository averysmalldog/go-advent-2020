package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

	// solution part 1
	numList := intList(list)
	diffby1 := 0
	diffby3 := 0
	sort.Ints(numList)
	startingJoltage := 0
	
	for k, v := range numList {
		difference := v-startingJoltage
		switch difference {
		case 1:{
			diffby1++
		}
		case 3:{
			diffby3++
		}
		default: {}
		}
		startingJoltage = v
		fmt.Printf("i: %d, v: %d, d1: %d, d3: %d\n", k, v, diffby1, diffby3)
	}
	diffby3++ // last jump
	fmt.Println(numList, startingJoltage, diffby1, diffby3, diffby1*diffby3)

}

// func listToMap(numList []string) (map[string]bool) {
// 	result := map[string]bool{}
// 	for _,v := range numList {
// 		result[v] = true
// 	}
// 	return result
// }

func intList(stringList []string) []int {
	intList := []int{}
	for _, v := range stringList {
		intValue, _ := strconv.Atoi(v)
		intList = append(intList, intValue)
	}
	return intList
}
