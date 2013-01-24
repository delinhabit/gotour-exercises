package main

import (
	"fmt"
	"runtime"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	visited := make(map[string]bool)

	var fetch func(string, int, chan bool)
	fetch = func(url string, depth int, done chan bool) {
		defer func() {
			done <- true
		}()

		if depth <= 0 || visited[url] {
			return
		}

		visited[url] = true
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("[%d] Found: %s %q\n", runtime.NumGoroutine(), url, body)
		for _, u := range urls {
			d := make(chan bool)
			go fetch(u, depth-1, d)
			<-d
		}
	}

	done := make(chan bool)
	go fetch(url, depth, done)
	<-done

	return
}

func main() {
	fmt.Printf("# Goroutines: %d\n", runtime.NumGoroutine())
	Crawl("http://golang.org/", 4, fetcher)
	fmt.Printf("# Goroutines: %d\n", runtime.NumGoroutine())
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("[%d] Not found: %s", runtime.NumGoroutine(), url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
