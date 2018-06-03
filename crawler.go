package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

type link struct {
	URL string
}

type topLinksSnapshot struct {
	Links   []link
	FoundAt time.Time
}

func main() {
	c := colly.NewCollector()
	snapshot := topLinksSnapshot{FoundAt: time.Now()}

	c.OnHTML(".storylink", func(e *colly.HTMLElement) {
		link := link{}
		link.URL = e.Attr("href")
		snapshot.Links = append(snapshot.Links, link)
	})

	c.Visit("https://news.ycombinator.com")
	fmt.Println(snapshot)
}
