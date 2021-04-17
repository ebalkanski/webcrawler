package crawler

import "sync"

type UrlsMap struct {
	mu   sync.Mutex
	urls map[string]int
}

func (urlsMap *UrlsMap) Visit(url string) {
	urlsMap.mu.Lock()
	urlsMap.urls[url] += 1
	urlsMap.mu.Unlock()
}

func (urlsMap *UrlsMap) IsVisited(url string) bool {
	urlsMap.mu.Lock()
	_, visited := urlsMap.urls[url]
	defer urlsMap.mu.Unlock()
	return visited
}
