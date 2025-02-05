package main

import (
	"errors"
	"fmt"
	"sync"
)

/*
need to understand

http://127.0.0.1:3999/concurrency/10
https://tour.go-zh.org/concurrency/10

http://127.0.0.1:3999/concurrency/11
https://tour.go-zh.org/concurrency/11

练习：Web 爬虫

在这个练习中，我们将会使用 Go 的并发特性来并行化一个 Web 爬虫。

修改 Crawl 函数来并行地抓取 URL，并且保证不重复。

提示：你可以用一个 map 来缓存已经获取的 URL，但是要注意 map 本身并不是并发安全的！

ref:
https://github.com/golang/tour/blob/master/solutions/webcrawler.go
*/

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// fetched tracks URLs that have been (or are being) fetched.
// The lock must be held while reading from or writing to the map.
// See https://golang.org/ref/spec#Struct_types section on embedded types
var fetched = struct {
	m map[string]error
	sync.Mutex
}{m:make(map[string]error)}

var loading = errors.New("url load in progress") // sentinel value

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	if depth <= 0 {
		fmt.Printf("<- Done with %v, depth 0.\n",url)
		return
	}

	fetched.Lock()
	if _,ok := fetched.m[url]; ok {
		fetched.Unlock()
		fmt.Printf("<- Done with %v, already fetched.\n", url)
		return
	}
	// We mark the url to be loading to avoid others reloading it at the same time.
	fetched.m[url] = loading
	fetched.Unlock()

	// We load it concurrently.
	body, urls, err := fetcher.Fetch(url)

	// And update the status in a synced zone.
	fetched.Lock()
	fetched.m[url] = err
	fetched.Unlock()

	if err != nil {
		//fmt.Println(err)
		fmt.Printf("<- Error on %v : %v\n", url, err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	done := make(chan bool)

	for i,u := range urls {
		fmt.Printf("-> Crawling child %v/%v of %v : %v.\n",i,len(urls),url,u)
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			done <- true
		}(u)
	}

	for i,u := range urls {
		fmt.Printf("<- [%v] %v/%v Waiting for child %v.\n", url, i, len(urls), u)
		<- done
	}
	fmt.Printf("<- Done with %v\n", url)
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher 是返回若干结果的 Fetcher。
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

// fetcher 是填充后的 fakeFetcher。 Mock Data
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

