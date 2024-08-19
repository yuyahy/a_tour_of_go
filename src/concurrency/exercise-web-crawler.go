package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu        sync.Mutex
	cache_map map[string]bool
}

func (c *SafeCounter) Update(key string) {
	// mutexでキャッシュ情報へのアクセスを排他制御
	c.mu.Lock()
	c.cache_map[key] = true
	c.mu.Unlock()
}

func (c *SafeCounter) Check(key string) bool {
	// mutexでキャッシュ情報へのアクセスを排他制御
	// ※ここはキャッシュを読み取っているだけなので、排他制御は不要かも
	c.mu.Lock()
	result := c.cache_map[key]
	c.mu.Unlock()

	return result
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// func Crawl(url string, depth int, fetcher Fetcher, cache_map map[string]bool) {
func Crawl(url string, depth int, fetcher Fetcher, safe_counter SafeCounter) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 || safe_counter.Check(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	// safe_counter.cache_map[url] = true
	safe_counter.Update(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	// URLのスライスに対して再帰的にクロールする処理を並列化しする
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			Crawl(url, depth-1, fetcher, safe_counter)
		}(u)
	}
	// 各goroutineの処理が終了するまで待機
	wg.Wait()

	return
}

func main() {
	//cache_map := make(map[string]bool)
	c := SafeCounter{cache_map: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, c)
}

// fakeFetcher is Fetcher that returns canned results.
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

// fetcher is a populated fakeFetcher.
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
