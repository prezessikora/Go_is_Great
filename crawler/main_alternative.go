package main

import (
	"fmt"
)

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl2(url string, depth int, fetcher Fetcher, urlsChan chan string) {
	if depth <= 0 {
		return
	}
	// visited guard
	if value := visited[url]; value {
		return
	} else {
		visited[url] = true
	}
	//visited guard
	body, urls, err := fetcher.Fetch(url)
	for _, newUrl := range urls {
		urlsChan <- newUrl
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

}

func main_alternative() {

	urlsChan := make(chan string)
	go func() {
		urlsChan <- "https://golang.org/"
	}()

	for {
		u := <-urlsChan
		go Crawl2(u, 4, fetcher, urlsChan)
	}

}
