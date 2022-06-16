package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("test", "test"),
	)
	c.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		DomainGlob: "test",
		// Set a delay between requests to these domains
		Delay: 1 * time.Second,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	})
	c.MaxDepth = 2
	c.AllowURLRevisit = true

	var inputText string
	contains := "test"
	var numTimesFound int
	//contains := os.Args[0]
	var LinksToVisit []string

	// Find and print all links
	c.OnHTML("body", func(e *colly.HTMLElement) {
		inputText = fmt.Sprintf(e.Text)
		numTimesFound = strings.Count(inputText, contains)

		fmt.Println(e.Text)
	})
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Term: ", contains, " found ", numTimesFound, " times on page", r.Request.URL)
	})

	c.OnHTML(".container", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")

		for _, j := range links {
			if strings.HasPrefix(j, "https://") == true {
				LinksToVisit = append(LinksToVisit, j)
			}
		}
		fmt.Println("future links:", LinksToVisit)

		//var visited []string

		/*
			for _, j := range LinksToVisit {
				for _, k := range visited {
					if k != j {
						err := c.Visit(j)
						if err != nil {
							fmt.Println("ERROR:", err)
						}
						visited = append(visited, j)
					}
				}
			}

		*/

		for _, j := range LinksToVisit {

			err := c.Visit(j)
			if err != nil {
				fmt.Println("ERROR:", err)

			}
		}

	})

	err := c.Visit("test")
	if err != nil {
		fmt.Println("ERROR:", err)
	}

}

func ContainsTextBool(input, search string) bool {
	foundText := strings.Contains(search, input)
	return foundText
}
