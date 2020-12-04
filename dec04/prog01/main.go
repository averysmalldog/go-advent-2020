package main

import (
	"io/ioutil"
	"fmt"
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
	fmt.Print(string(dat))

	// solution
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	separator := "\n\n"
	validPassports := 0
	entries := strings.Split(string(dat), separator)

	for k, v := range entries {
		passportEntry := make(map[string]string)
		fields := strings.Fields(v)
		for _, v1 := range fields {
			tuple := strings.Split(v1, ":")
			passportEntry[tuple[0]] = tuple[1]
			fmt.Printf("Entry %d k: %s v: %s.\n", k, tuple[0], passportEntry[tuple[0]])
		}
		missingData := 0
		for _, v1 := range requiredFields {
			if passportEntry[v1] == "" {
				fmt.Printf("Passport %d is missing %s.\n", k, v1)
				missingData++
			} 
		}
		if missingData == 0 {
			fmt.Printf("Passport %d is VALID.\n", k)
			validPassports++
		}
		fmt.Println()
	}
	fmt.Printf("Total Passports: %d.\n", len(entries))
	fmt.Printf("Valid Passports: %d.\n", validPassports)
}
