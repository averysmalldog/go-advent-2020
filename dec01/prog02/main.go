package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"time"
)

func getInput(day int) {
	dayString := strconv.Itoa(day)
	myInputURL := "https://adventofcode.com/2020/day/"+dayString+"/input"

	// check for the session ID
	key := os.Getenv("ADVENT_SESSION")

	// it's good practice!
	timeout := time.Duration(5 * time.Second)

	// get your auth cookie set up
	jar, _ := cookiejar.New(nil)
	var cookies []*http.Cookie
	cookie := &http.Cookie{
		Name:   "session",
		Value:  key,
		Path:   "/",
		Domain: ".adventofcode.com",
	}
	cookies = append(cookies, cookie)
	u, _ := url.Parse(myInputURL)
	jar.SetCookies(u, cookies)

	// set up the client with that jar
	client := http.Client{
		Timeout: timeout,
		Jar: jar,
	}

	// get the data
	resp, err := client.Get(myInputURL)
	if err != nil {
		panic("request failed")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))
}

func findAddsTo2020(list []int) (int, int, int) {
	for i1, n1 := range list {
		for i2, n2 := range list[i1+1:] {
			for i3, n3 :=range list[i2+1:]{
				if(n1+n2+n3 == 2020){
					fmt.Printf("Found a match!\n\tIndex: %d, Value: %d\n\tIndex: %d, Value %d\n\tIndex: %d, Value %d\n",i1, n1, i2, n2, i3, n3)
					return n1, n2, n3
				}
			}
		}
	}
	return 0, 0, 0
}

func getSecretKey(int1 int, int2 int, int3 int) int {
	return int1*int2*int3
}

func main() {
	getInput(1)
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
	int1, int2, int3 := findAddsTo2020(list)
	secretKey := getSecretKey(int1, int2, int3)
	fmt.Printf("Secret Key: %d\n", secretKey)
}