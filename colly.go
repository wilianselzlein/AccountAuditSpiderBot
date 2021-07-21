package main

import (
	"fmt"
	"log"
	"time"
	"github.com/gocolly/colly"
	"AccountAuditSpiderBot/spider"
)

func init() {
	fmt.Printf("Yaml: %+v\n", spider.LoadConfig())
}

func test() {
	// create a new collector
	c := colly.NewCollector(
		// MaxDepth is 1, so only the links on the scraped page
		// is visited, and no further links are followed
		colly.MaxDepth(1),
	)

	// authenticate
	err := c.Post(spider.Y.Url[0].Login, 
		map[string]string{
			"username": spider.Y.User[0].Login,
			"password": spider.Y.User[0].Password,
			"submit": "Entrar"})
	if err != nil {
		log.Fatal(err)
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		log.Println("+HTML+")
		log.Println(e.Text)
		time.Sleep(1 * time.Second)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visit to", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	//TODO: Copy cookies from selenium
	// start scraping
	// HTTP Post Binding (Request)<P>JavaScript is disabled. We strongly recommend to enable it. Click the button below to continue.
	// </P><INPUT TYPE="SUBMIT" VALUE="CONTINUE" />

	c.Visit(spider.Y.Url[0].List_Accounts)

}
