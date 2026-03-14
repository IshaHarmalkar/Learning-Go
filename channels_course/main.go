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

	c := make(chan string)


	//we only use the go key word in front of a fn call.

	for _, link := range links {
		go checkLink(link, c)

	}

	//infinite loop, l -> fn literal
	for l := range c {

		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}


	//a value coming from chan is a blocking call

	//most of the time the main rouine should not be put to sleep


   //In practice, we never ever try to reference the same variable inside of
   //two different routines
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