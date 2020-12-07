package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	Adjective string
	Color string
	Rules []Rule
}

type Rule struct {
	Quantity int
	Bag Bag
}

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
	bagList := make(map[string]map[string][]Rule)
	for k1, rule := range list {
		ruleset := strings.Split(rule, "contain ")

		// define the bag the rule is for
		rootBagStrings := strings.Fields(ruleset[0])
		rootBag := Bag{
			Adjective: rootBagStrings[0],
			Color: rootBagStrings[1],
		}
		fmt.Printf("%+v\n",rootBag)
		
		// set up the rules for the bag
		thisBagRules := []Rule{}
		items := strings.Split(ruleset[1], ", ")
		for _, item := range items {
			elements := strings.Fields(item)
			numBags, err := strconv.Atoi(elements[0])
			if err != nil {
				fmt.Printf("\terror converting to int.")
			}
			thisRule := Rule{
				Quantity: numBags,
				Bag: Bag{
					Adjective: elements[1],
					Color: elements[2],
				},
			}
			thisBagRules = append(thisBagRules, thisRule)
		}
		
		// compose
		theBag := Bag{
			Adjective: rootBag.Adjective,
			Color: rootBag.Color,
			Rules: thisBagRules,
		}
		fmt.Printf("Rule %d details:\n\tA %s %s bag contains:\n", k1, theBag.Adjective, theBag.Color)
		for _, v := range theBag.Rules{
			fmt.Printf("\t%d %s %s\n", v.Quantity, v.Bag.Adjective, v.Bag.Color)
		}
		bagList[theBag.Adjective][theBag.Color] = thisBagRules
	}
	fmt.Printf("%+v", bagList)
}