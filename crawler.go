package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

type Story struct {
	Rank  string
	Title string
	URL   string
}

type topStoriesSnapshot struct {
	Stories []Story
	FoundAt time.Time
}

func main() {
	c := colly.NewCollector()
	snapshot := topStoriesSnapshot{FoundAt: time.Now()}

	c.OnHTML(".athing", func(e *colly.HTMLElement) {
		story := CreateStory(e)
		snapshot.Stories = append(snapshot.Stories, story)
	})

	c.Visit("https://news.ycombinator.com")
	fmt.Println(snapshot)
}

func CreateStory(e *colly.HTMLElement) Story {
	rank := e.ChildText(".rank")
	url := e.ChildAttr(".storylink", "href")
	title := e.ChildText(".storylink")

	return Story{Rank: rank, URL: url, Title: title}
}
