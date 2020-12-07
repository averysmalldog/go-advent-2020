package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Quantity int
	Bag string
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
	bagList := make(map[string][]Rule)
	for _, rule := range list {
		ruleset := strings.Split(rule, "contain ")

		// define the bag the rule is for
		rootBagStrings := strings.Fields(ruleset[0])
		rootBag := fmt.Sprintf("%s %s", rootBagStrings[0],rootBagStrings[1])
		//fmt.Printf("%+v\n",rootBag)
		
		// set up the rules for the bag
		thisBagRules := []Rule{}
		if ruleset[1] != "no other bags." {
			items := strings.Split(ruleset[1], ", ")
			for _, item := range items {
				elements := strings.Fields(item)
				numBags, err := strconv.Atoi(elements[0])
				if err != nil {
					//fmt.Printf("\terror converting to int.")
				}
				thisRule := Rule{
					Quantity: numBags,
					Bag: fmt.Sprintf("%s %s", elements[1], elements[2]),
				}
				thisBagRules = append(thisBagRules, thisRule)
			}
		
		// compose
		bagList[rootBag] = thisBagRules
		}
	}
	containsList := map[string]int{}
	walkDown(bagList, &containsList, "shiny gold")
	fmt.Printf("Number of bag colors contained by shiny gold: %d\n", len(containsList))
}

func walkUp(ruleMap map[string][]Rule, containsList *map[string]bool, desiredBag string) {
	list := *containsList
	for k, rulelist := range ruleMap {
		for _, rule := range rulelist{
			//fmt.Printf("Rule %s, Item %d: %+v\n", k, k1, rule)
			if rule.Bag == desiredBag{
				list[k] = true
				fmt.Printf("%s contains %s.\n", k, desiredBag)
				walkUp(ruleMap, &list, k)
			}
		}
	}
}

// need to keep track of rules with 0 to solve this
func walkDown(ruleMap map[string][]Rule, containsList *map[string]int, desiredBag string) {
	list := *containsList
	for k, rulelist := range ruleMap {
		if k == desiredBag{
			for _, rule := range rulelist{
				list[k] += rule.Quantity
				fmt.Printf("%s contains %d %s.\n", desiredBag, rule.Quantity, rule.Bag)
				walkDown(ruleMap, &list, k)
			}
		}
	}
}