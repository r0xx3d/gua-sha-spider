package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/http/cookiejar"
	"time"
)

var visitedurls = make(map[string]bool)

func main() {
	seedurl := "https:/"+"/crawler-test.com/"

	crawl(seedurl, 0)
}

func crawl(currenturl string, maxdepth int) {

	c := colly.NewCollector(
		colly.AllowedDomains("crawler-test.com"),
		colly.MaxDepth(maxdepth),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob: "*",
		Parallelism: 5,
		Delay: 2 * time.Second,
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" && !visitedurls[link]{
			visitedurls[link] = true
			fmt.Println(link)
			e.Request.Visit(link)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (compatible; Colly/2.1; +https:/"+"/github.com/gocolly/colly")
		fmt.Println("Crawling", r.URL)
	})


	cookiesJar,_ := cookiejar.New(nil)
	c.SetCookieJar(cookiesJar)


	c.OnError(func(e *colly.Response, err error) {
		fmt.Println("Request URL:", e.Request.URL, "failed with response",e,"\nError:",err)
	})

	err := c.Visit(currenturl)
	if err != nil {
		fmt.Println("Error visiting page:", err)
	}

	c.Wait()

}
