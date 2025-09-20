package main

import (
	"fmt"
	"net/http"
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

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}


	//a value coming from chan is a blocking call



}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")

		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")

	c <- "Yup it's up"
}