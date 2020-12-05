package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// validation functions
func validateByr(data string) bool {
	fmt.Printf("%s", data)
	if len(data) != 4 {
		return false
	}
	dob, err := strconv.Atoi(data)
	if err != nil {
		return false
	}
	if dob >= 1920 && dob <= 2002 {
		return true
	}
	return false
}

func validateIyr(data string) bool {
	fmt.Printf("%s", data)
	if len(data) != 4 {
		return false
	}
	dob, err := strconv.Atoi(data)
	if err != nil {
		return false
	}
	if dob >= 2010 && dob <= 2020 {
		return true
	}
	return false
}

func validateEyr(data string) bool {
	fmt.Printf("%s", data)
	if len(data) != 4 {
		return false
	}
	dob, err := strconv.Atoi(data)
	if err != nil {
		return false
	}
	if dob >= 2020 && dob <= 2030 {
		return true
	}
	return false
}

func validateHgt(data string) bool {
	fmt.Printf("%s", data)
	if strings.Contains(data, "cm") {
		digits := strings.Split(data, "c")
		height, err := strconv.Atoi(digits[0])
		if err != nil {
			return false
		}
		if height >= 150 && height <= 193 {
			return true
		}
	} else if strings.Contains(data, "in") {
		digits := strings.Split(data, "i")
		height, err := strconv.Atoi(digits[0])
		if err != nil {
			return false
		}
		if height >= 59 && height <= 76 {
			return true
		}
	}
	return false
}

func validateHcl(data string) bool {
	fmt.Printf("%s", data)
	match, err := regexp.Match(`#[0-9a-f]{6}`, []byte(data))
	if err != nil {
		return false
	}
	if match {
		return true
	}
	return false
}

func validateEcl(data string) bool {
	fmt.Printf("%s", data)
	validEcls := map[string]int{
		"amb": 1,
		"blu": 1,
		"brn": 1,
		"gry": 1,
		"grn": 1,
		"hzl": 1,
		"oth": 1,
	}
	if validEcls[data] == 1 {
		return true
	}
	return false
}

func validatePid(data string) bool {
	fmt.Printf("%s", data)
	if len(data) != 9 {
		return false
	}
	match, err := regexp.Match(`[0-9]{9}`, []byte(data))
	if err != nil {
		return false
	}
	if match {
		return true
	}
	return false
}

func validateCid(data string) bool {
	return true
}

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
	output := []string{}

	for k, v := range entries {
		passportEntry := make(map[string]string)
		fields := strings.Fields(v)
		for _, v1 := range fields {
			tuple := strings.Split(v1, ":")
			passportEntry[tuple[0]] = tuple[1]
			fmt.Printf("Entry %d k: %s v: %s.\n", k, tuple[0], passportEntry[tuple[0]])
		}
		missingData := 0
		invalidData := 0
		for _, v1 := range requiredFields {
			switch v1 {
			case "byr":
				{
					fmt.Printf("\tChecking byr: ")
					if validateByr(passportEntry[v1]) {
						fmt.Printf(" valid.\n")
					} else {
						fmt.Printf(" INVALID.\n")
						invalidData++
					}
				}
			case "iyr":
				{
					fmt.Printf("\tChecking iyr: ")
					if validateIyr(passportEntry[v1]) {
						fmt.Printf(" valid.\n")
					} else {
						fmt.Printf(" INVALID.\n")
						invalidData++
					}
				}
			case "eyr":
				{
					fmt.Printf("\tChecking eyr: ")
					if validateEyr(passportEntry[v1]) {
						fmt.Printf(" valid.\n")
					} else {
						fmt.Printf(" INVALID.\n")
						invalidData++
					}
				}
			case "hgt":
				{
					fmt.Printf("\tChecking hgt: ")
					if validateHgt(passportEntry[v1]) {
						fmt.Printf(" valid.\n")
					} else {
						fmt.Printf(" INVALID.\n")
						invalidData++
					}
				}
			case "hcl":
				{
					fmt.Printf("\tChecking hcl: ")
					if validateHcl(passportEntry[v1]) {
						fmt.Printf(" valid.\n")
					} else {
						fmt.Printf(" INVALID.\n")
						invalidData++
					}
				}
			case "ecl":
				{
					fmt.Printf("\tChecking ecl: ")
					if validateEcl(passportEntry[v1]) {
						fmt.Printf(" valid.\n")
					} else {
						fmt.Printf(" INVALID.\n")
						invalidData++
					}
				}
			case "pid":
				{
					fmt.Printf("\tChecking pid: ")
					if validatePid(passportEntry[v1]) {
						fmt.Printf(" valid.\n")
					} else {
						fmt.Printf(" INVALID.\n")
						invalidData++
					}
				}
			}
			if passportEntry[v1] == "" {
				fmt.Printf("\tMissing %s.\n", v1)
				missingData++
			}
		}
		if missingData == 0 && invalidData == 0 {
			fmt.Printf("%d\n", k)
			validPassports++
		}
		fmt.Println()
		output = append(output, fmt.Sprintf("%d,%s,%s,%s,%s,%s,%s,%s", k, passportEntry["byr"], passportEntry["iyr"], passportEntry["eyr"], passportEntry["hgt"], passportEntry["hcl"], passportEntry["ecl"], passportEntry["pid"]))
	}
	fmt.Printf("Total Passports: %d.\n", len(entries))
	fmt.Printf("Valid Passports: %d.\n", validPassports)
	// for _,value := range output{
	// 	fmt.Println(value)
	// }

}
