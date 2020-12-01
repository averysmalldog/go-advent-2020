package main

import "fmt"

func findAddsTo2020(list []int) (int, int) {
	return 1, 2
}

func getSecretKey(int1 int, int2 int) int {
	return 3
}

func main() {
	fakeList := []int{2019, 1, 50}
	int1, int2 := findAddsTo2020(fakeList)
	secretKey := getSecretKey(int1, int2)
	fmt.Printf("Integer #1: %d\nInteger #2: %d\n",
		int1, 
		int2)
	fmt.Printf("Secret Key: %d\n", secretKey)
}