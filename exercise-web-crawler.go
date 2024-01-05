package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type SafeCache struct {
	mu    sync.Mutex
	cache map[string]bool
}

func NewSafeCache() *SafeCache {
	return &SafeCache{
		cache: make(map[string]bool),
	}
}

func Crawl(url string, depth int, fetcher Fetcher, cache *SafeCache, ch chan string) {
	cache.mu.Lock()

	if cache.cache[url] {
		cache.mu.Unlock()
		ch <- url
		return
	}

	cache.cache[url] = true
	cache.mu.Unlock()

	if depth <= 0 {
		ch <- url
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		ch <- url
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	urlCh := make(chan string)

	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, cache, urlCh)
	}

	for range urls {
		<-urlCh
	}

	ch <- url
}

func main() {
	cache := NewSafeCache()
	ch := make(chan string)

	go Crawl("https://golang.org/", 4, fetcher, cache, ch)

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	close(ch)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
