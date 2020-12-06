package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// test for presence of input.txt
	_, err := os.Stat("input.txt")
	if os.IsNotExist(err) {
		fmt.Println("input.txt not found. Exiting program.")
		return
	}

	// read entire file as bytes
	path := "input.txt"
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	// solution
	separator := "\n\n"
	totalYeses := 0
	entries := strings.Split(string(dat), separator)
	for k, v := range entries {
		answerMap := map[rune]int{}
		answerSum := 0
		fields := strings.Fields(v)
		for k1, v1 := range fields {
			for _, v2 := range v1 {
				answerMap[v2] += 1
			}
			fmt.Printf("Group %d Entry %d: %s, size of group: %d\n", k, k1, v1, len(fields))
		}
		for k2, i := range answerMap {
			fmt.Printf("\tk:%c i:%d\n", k2, i)
			if i == len(fields){
				answerSum += 1
			}
			
		}
		fmt.Printf("Group %d total YESes: %d\n\n", k, answerSum)
		totalYeses += answerSum
	}
	fmt.Printf("Sum of all answers: %d\n", totalYeses)

}