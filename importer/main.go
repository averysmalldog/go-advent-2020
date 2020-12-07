// README: you can find your session string at chrome://settings/cookies/detail?site=adventofcode.com&search=cookies
// README: set your env var as follows:
// export ADVENT_SESSION="<your-session-string>"
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"time"
)

// getInput is generalized for all Advent of Code programs.
func getInput(day int) {
	dayString := strconv.Itoa(day)
	myInputURL := "https://adventofcode.com/2020/day/" + dayString + "/input"

	// check for the session ID
	key, exists := os.LookupEnv("ADVENT_SESSION")
	if !exists {
		fmt.Println("WARNING! Environment variable ADVENT_SESSION is not set.")
		fmt.Println("Skipping HTTPs call for your personalized inputs.")
		return
	}

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
		Jar:     jar,
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

	// now that you have the data, write it to input.txt
	file, err := os.Create("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	numBytes, err := file.Write(body)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	fmt.Println("input.txt created and populated successfully. ", numBytes, " bytes written.")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	// make and populate the file if it isn't there
	_, err := os.Stat("input.txt")
	if os.IsNotExist(err) {
		fmt.Println("input.txt not found. Making HTTPs call to generate it.")
		getInput(7)
	}
}
