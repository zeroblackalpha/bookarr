package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.yuanzunxsw.com"),
		colly.MaxDepth(1),
		colly.DetectCharset(),
	)

	c.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		DomainGlob: "www.yuanzunxsw.com/*",
		// Set a delay between requests to these domains
		Delay: 1 * time.Second,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	})

	numVisited := 0
	c.OnRequest(func(r *colly.Request) {
		if numVisited > 3000 {
			r.Abort()
		}
		numVisited++
	})

	// On every a element which has href attribute call callback
	c.OnHTML("div#chapterlist a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// On every a element which has href attribute call callback
	c.OnHTML("div.title > h1", func(e *colly.HTMLElement) {
		fmt.Println(strings.ReplaceAll(e.Text, "正文 ", ""))
		fmt.Println()
	})

	// On every a element which has href attribute call callback
	c.OnHTML("div#content", func(e *colly.HTMLElement) {
		fmt.Println(strings.ReplaceAll(e.Text, "\n\n", "\n"))
		fmt.Println()
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.yuanzunxsw.com/txt/64019/")
}
