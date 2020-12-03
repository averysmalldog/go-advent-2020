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
	validCounter := 0

	for k, v := range list {
		fmt.Printf("Entry %d: %s: ", k, v)
		// Parse string into fields
		str := strings.Fields(v)

		minMax := strings.Split(str[0], "-")
		index1, _ := strconv.Atoi(minMax[0])
		index2, _ := strconv.Atoi(minMax[1])
		desiredLetter := string(str[1][0])
		password := str[2]

		// get count if match
		if string(password[index1-1]) == desiredLetter {
			if string(password[index2-1]) == desiredLetter {
				fmt.Printf("not valid\n")
			} else {
				validCounter++
				fmt.Printf("valid!\n")
			}
		} else {
			if string(password[index2-1]) == desiredLetter {
				validCounter++
				fmt.Printf("valid!\n")
			} else {
				fmt.Printf("not valid\n")
			}
		}
	}

	fmt.Printf("Valid passwords: %d\n", validCounter)

}
