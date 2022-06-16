package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	var inputText string
	contains := "web scraping"
	//contains := os.Args[0]

	// Find and print all links
	c.OnHTML("body", func(e *colly.HTMLElement) {
		inputText = fmt.Sprintf(e.Text)
		output := strings.Count(inputText, contains)
		fmt.Println(e.Name, output)
	})

	err := c.Visit("https://en.wikipedia.org/wiki/Web_scraping")
	if err != nil {
		fmt.Println(err)
		return
	}

}

func ContainsTextBool(input, search string) bool {
	foundText := strings.Contains(search, input)
	return foundText
}
