package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
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

	// SOLUTION
	// desiredLetter := "X"
	positionX := 0
	positionY := 0
	movementX := 3
	movementY := 1
	treesHit  := 0
	var positionXMod int


	for k, v := range list {
		if (k == positionY){
			positionXMod = len(v)
			truePositionX := positionX % positionXMod

			if (string(v[truePositionX]) == "#") {
				treesHit++
				if (truePositionX == 0) {
					fmt.Printf("Line %d: %s%s\n", k, "X", v[truePositionX+1:])
				} else {
					fmt.Printf("Line %d: %s%s%s\n", k, v[0:truePositionX-1], "X", v[truePositionX+1:])
				}
			}
			if (string(v[truePositionX]) == ".") {
				if (truePositionX == 0) {
					fmt.Printf("Line %d: %s%s\n", k, "O", v[truePositionX+1:])
				} else {
					fmt.Printf("Line %d: %s%s%s\n", k, v[0:truePositionX-1], "O", v[truePositionX+1:])
				}
			}

			positionX += movementX
			positionY += movementY
		}
	}

	fmt.Printf("Trees hit: %d\n", treesHit)
}
