package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("hello")
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

}
