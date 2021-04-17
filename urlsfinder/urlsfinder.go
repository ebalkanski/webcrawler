package urlsfinder

import (
	"fmt"
	"math/rand"
	"time"
)

type UrlsFinder interface {
	// Find returns all urls which are present on the page(Url) which is passed as argument
	FindAll(url string) (urls []string)
}

type FakeUrlsFinder struct {
	urls []string
}

func (finder *FakeUrlsFinder) Init() {
	urlsCount := 100
	finder.urls = make([]string, urlsCount)
	for i := 0; i < urlsCount; i++ {
		finder.urls[i] = fmt.Sprintf("www.go.org %v", i)
	}
}

func (finder *FakeUrlsFinder) FindAll(url string) []string {
	rand.Seed(time.Now().UnixNano())
	sleep := rand.Intn(5)
	time.Sleep(time.Duration(sleep) * time.Second)
	return finder.urls
}
