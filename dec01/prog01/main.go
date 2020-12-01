package main

import (
	"fmt"
)

func findAddsTo2020(list []int) (int, int) {
	// var int1, int2 int
	for i1, n1 := range list {
		fmt.Printf("range: index: %d, value: %d\n", i1, n1)
		for _, n2 := range list[i1+1:] {
			fmt.Printf("\tsum of %d and %d is %d\n", n1, n2, n1+n2)
			if(n1+n2 == 2020){
				return n1, n2
			}
		}
	}
	return 0, 0
}

func getSecretKey(int1 int, int2 int) int {
	return 3
}

func main() {
	fakeList := []int{2019, 50, 40, 30, 20, 10, 1}
	int1, int2 := findAddsTo2020(fakeList)
	secretKey := getSecretKey(int1, int2)
	fmt.Printf("Integer #1: %d\nInteger #2: %d\n",
		int1, 
		int2)
	fmt.Printf("Secret Key: %d\n", secretKey)
}