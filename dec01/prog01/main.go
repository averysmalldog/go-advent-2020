package main

import "fmt"

func findAddsTo2020(list []int) (int, int){
	return 1, 2
}

func main() {
	fakeList := []int{1, 2}
	int1, int2 := findAddsTo2020(fakeList)
	fmt.Printf("Integer #1: %d.\nInteger #2: %d.\n",
		int1, 
		int2)
}