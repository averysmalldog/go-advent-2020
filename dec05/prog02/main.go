package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func convertToInt(one string, input string) (int, error) {
	value := 0
	for k, v := range input {
		if v == rune(one[0]) {
			value += int(math.Pow(2, float64(len(input)-k-1)))
		}
	}
	return value, nil
}

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
	highestSid := 0
	sidList := []int{}
	for k, v := range list {
		row, _ := convertToInt("B", v[:7])
		column, _ := convertToInt("R", v[7:])
		sid := row*8 + column
		sidList = append(sidList, sid)
		if sid > highestSid {
			highestSid = sid
		}

		fmt.Printf("Entry %d: %s, Row: %d, Column: %d, Seat ID: %d\n", k, v, row, column, sid)
	}
	fmt.Printf("Highest Seat ID: %d\n", highestSid)
	sort.Ints(sidList)
	for k, _ := range sidList {
		if k > 0 && k < len(sidList) {
			if sidList[k-1] != sidList[k]-1 {
				fmt.Printf("Missing Seat ID: %d\n", sidList[k]-1)
			}
		}
	}
}
