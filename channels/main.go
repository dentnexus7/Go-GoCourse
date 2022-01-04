package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Create a new channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// First example
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// Repeating routines example
	// for {
	// 	checkLink(<-c, c)
	// }

	// Alternative Loop Syntax Example
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
