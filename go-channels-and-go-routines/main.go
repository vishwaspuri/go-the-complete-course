package main

import (
	"fmt"
	"net/http"
	"time"
)

// Main Routine
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// Creating a new go routine
		go checkLink(link, c)
	}
	// Creating c esque for loops
	// for i := 0; i < len(links); i++ {
	// 	checkLink(<-c, c)
	// }

	// Creating infinite loop
	for l := range c { // Means wait for a channel to return a value
		// Not very clever
		// time.Sleep(time.Second)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, error := http.Get(link)
	if error != nil {
		fmt.Println(link + " might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
