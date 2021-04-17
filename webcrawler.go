package main

import (
	"fmt"
	"go/webcrawler/crawler"
	"time"
)

func main() {
	done := make(chan bool)
	crawler := crawler.WebCrawler{}

	timeStart := time.Now()
	go crawler.Start("www.go.org", done)
	<-done
	timeDiff := time.Since(timeStart)
	fmt.Printf("\nElapsed time %.2f sec.\n\n", timeDiff.Seconds())
}
