package crawler

import (
	"fmt"
	"go/webcrawler/urlsfinder"
	"sync"
	"time"
)

type Crawler interface {
	Start(url string, done chan<- bool)
}

type WebCrawler struct{}

func (crawler WebCrawler) Start(url string, done chan<- bool) {
	fmt.Println("\nStart crawling...\n")
	time.Sleep(time.Second)

	depth := 5
	fakeUrlsFinder := &urlsfinder.FakeUrlsFinder{}
	fakeUrlsFinder.Init()
	visitedUrls := &UrlsMap{urls: make(map[string]int)}
	buffer := 0
	watcher := make(chan string, buffer)

	// Watch working go routines in console
	go crawler.watch(watcher)

	// Start the crawler and wait for all cascade jobs to finish
	var wg sync.WaitGroup
	wg.Add(1)
	go crawler.crawl(url, visitedUrls, fakeUrlsFinder, depth, &wg, watcher)
	wg.Wait()

	// Print how many times each url has been visited and finish
	crawler.report(visitedUrls)

	done <- true
}

func (crawler WebCrawler) watch(watcher chan string) {
	for {
		fmt.Println(<-watcher)
	}
}

func (crawler WebCrawler) crawl(url string, visitedUrls *UrlsMap, urlsFinder urlsfinder.UrlsFinder, depth int, wg *sync.WaitGroup, watcher chan string) {
	defer wg.Done()

	if depth <= 0 {
		return
	}
	if !visitedUrls.IsVisited(url) {
		watcher <- fmt.Sprintf("Process url %v from depth %v", url, depth)
		urls := urlsFinder.FindAll(url)
		visitedUrls.Visit(url)
		watcher <- fmt.Sprintf("Processed url %v from depth %v", url, depth)
		for _, url := range urls {
			wg.Add(1)
			go crawler.crawl(url, visitedUrls, urlsFinder, depth-1, wg, watcher)
		}
	}
}

func (crawler WebCrawler) report(visitedUrls *UrlsMap) {
	fmt.Println()

	visitedMoreThanOnce := 0
	for url, count := range visitedUrls.urls {
		if count > 1 {
			fmt.Printf("%v is visited %v time(s)\n", url, count)
			visitedMoreThanOnce++
		}
	}
	if visitedMoreThanOnce > 0 {
		fmt.Printf("\nUrls visited more than once: %v\n", visitedMoreThanOnce)
	}
}
