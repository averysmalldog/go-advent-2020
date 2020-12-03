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

	// SOLUTION
	treesHit1 := findTreeHits(1, 1, list)
	treesHit2 := findTreeHits(3, 1, list)
	treesHit3 := findTreeHits(5, 1, list)
	treesHit4 := findTreeHits(7, 1, list)
	treesHit5 := findTreeHits(1, 2, list)
	fmt.Printf("Trees hit in run 1: %d\n", treesHit1)
	fmt.Printf("Trees hit in run 2: %d\n", treesHit2)
	fmt.Printf("Trees hit in run 3: %d\n", treesHit3)
	fmt.Printf("Trees hit in run 4: %d\n", treesHit4)
	fmt.Printf("Trees hit in run 5: %d\n", treesHit5)
	fmt.Printf("Multiplied Trees: %d\n", treesHit1*treesHit2*treesHit3*treesHit4*treesHit5)

}

func findTreeHits(movementX int, movementY int, data []string) int {
	positionX := 0
	positionY := 0
	treesHit := 0
	var positionXMod int

	for k, v := range data {
		if k == positionY {
			positionXMod = len(v)
			truePositionX := positionX % positionXMod

			if string(v[truePositionX]) == "#" {
				treesHit++
				if truePositionX == 0 {
					fmt.Printf("Line %d: %s%s\n", k, "X", v[truePositionX+1:])
				} else if truePositionX == positionXMod-1 {
					fmt.Printf("Line %d: %s%s\n", k, v[0:truePositionX], "X")
				} else {
					fmt.Printf("Line %d: %s%s%s\n", k, v[0:truePositionX], "X", v[truePositionX+1:])
				}
			}
			if string(v[truePositionX]) == "." {
				if truePositionX == 0 {
					fmt.Printf("Line %d: %s%s\n", k, "O", v[truePositionX+1:])
				} else if truePositionX == positionXMod-1 {
					fmt.Printf("Line %d: %s%s\n", k, v[0:truePositionX], "O")
				} else {
					fmt.Printf("Line %d: %s%s%s\n", k, v[0:truePositionX], "O", v[truePositionX+1:])
				}
			}

			positionX += movementX
			positionY += movementY
		}
	}

	fmt.Printf("Trees hit: %d\n", treesHit)
	return treesHit
}
