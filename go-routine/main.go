package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// define links slice with the type of string
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create a channel with the data type of string
	c := make(chan string)

	// loop over the links slice
	// each link will be input of the function checkLink
	// each go-routine will take care for each time function
	// is called.
	for _, link := range links {
		go checkLink(link, c)
	}

	// continue run checklink
	// after finishing check the input link
	// Waiting 5s, the go routine will check that link again

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// check whether the link (url) is up or not
// using http.Get(link)
// Note that channel is one of the arguments of function
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		// Send result to channel
		c <- link
		return
	}
	fmt.Println(link + " is up!")
	// Send result to channel
	c <- link
}
