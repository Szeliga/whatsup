package main

import (
	"fmt"
	"os"

	"github.com/SlyMarbo/rss"
)

func fetchFeed(url string) (*rss.Feed, error) {
	return rss.Fetch("http://feeds.feedburner.com/Antyweb?format=xml")
}

func main() {
	feed, err := fetchFeed("http://feeds.feedburner.com/Antyweb?format=xml")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(feed.Title)
}
