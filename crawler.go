package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	links := []string{}

	c.OnHTML(".storylink", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		links = append(links, link)
	})

	c.Visit("https://news.ycombinator.com")
	fmt.Println(links)
}
