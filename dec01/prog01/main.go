package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func findAddsTo2020(list []int) (int, int) {
	for i1, n1 := range list {
		for i2, n2 := range list[i1+1:] {
			if n1+n2 == 2020 {
				fmt.Printf("Found a match!\n\tIndex: %d, Value: %d\n\tIndex: %d, Value %d\n", i1, n1, i2, n2)
				return n1, n2
			}
		}
	}
	return 0, 0
}

func getSecretKey(int1 int, int2 int) int {
	return int1 * int2
}

func main() {
	// read input from the file!
	path := "input.txt"
	var list []int

	buf, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	snl := bufio.NewScanner(buf)
	for snl.Scan() {
		num, err := strconv.Atoi(snl.Text())
		if err != nil {
			log.Println("well I guess there's an error")
		}
		list = append(list, num)
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}

	// My solution
	// fakeList := []int{2019, 50, 40, 30, 20, 10, 1}
	int1, int2 := findAddsTo2020(list)
	secretKey := getSecretKey(int1, int2)
	fmt.Printf("Secret Key: %d\n", secretKey)
}
