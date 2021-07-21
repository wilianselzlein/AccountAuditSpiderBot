package main

import (
	"log"
	"time"
	"github.com/gocolly/colly"
	"AccountAuditSpiderBot/spider"
)

func init() {
	ftm.Printf("Yaml: %+v\n", config.load())
}

func main() {
	// create a new collector
	c := colly.NewCollector(
		// MaxDepth is 1, so only the links on the scraped page
		// is visited, and no further links are followed
		colly.MaxDepth(1),
	)

	// authenticate
	err := c.Post(y.Url[0].Login, 
		map[string]string{"username": y.User[0].Login, "password": y.User[0].Password, "submit": "Entrar"})
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

	c.Visit(y.Url[0].List_Accounts)

}
