package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	// "strings"
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
	invalidNumber := 0
	for k, v := range list[25:] {
		cache := listToMap(list[k:k+25])
		value, _ := strconv.Atoi(v)
		if !validate(cache, value) {
			fmt.Printf("Entry %d is invalid. Value: %s\n",k+25, v)
			invalidNumber = strconv.Atoi(v)
		}
	}

	// solution part 2
	magicNumberList := findMagicNumberList(list, invalidNumber)
	magicNumber := sumMinMax(magicNumberList)
}

func listToMap(numList []string) (map[string]bool) {
	result := map[string]bool{}
	for _,v := range numList {
		result[v] = true
	}
	return result
}

func validate(cache map[string]bool, value int) bool {
	for k, _ := range cache {
		factor, _ := strconv.Atoi(k)
		complement := value-factor
		if cache[strconv.Itoa(complement)] {return true}
	}
	return false
}

func findMagicNumberList(list []string, invalidNumber int) []int {
	
}