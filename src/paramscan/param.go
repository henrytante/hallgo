package paramscan

import (
	"fmt"
	"log"
	"net/url"

	cl "github.com/fatih/color"
	"github.com/gocolly/colly"
)

func ParamScan() {
	var green = cl.New(cl.FgHiGreen, cl.Bold)
	red := cl.New(cl.FgRed, cl.Bold)
	var url string
	red.Print("Digite a sua url (http://exemple.com/): ")
	fmt.Scanln(&url)
	urlToScan := url
	c := colly.NewCollector()

	// Slice para armazenar query strings únicas
	uniqueQueryStrings := make([]string, 0)

	// Callback para extrair query strings das URLs
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		queryString := extractQueryStringFromURL(link)
		if queryString != "" && !contains(uniqueQueryStrings, queryString) {
			uniqueQueryStrings = append(uniqueQueryStrings, queryString)
			green.Printf("Parâmetro encontrado: %v%s\n",url, queryString)
		}
	})

	err := c.Visit(urlToScan)
	if err != nil {
		log.Fatal(err)
	}
}

func extractQueryStringFromURL(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		return ""
	}

	query := u.RawQuery
	if query != "" {
		return "?" + query
	}

	return ""
}

func contains(slice []string, query string) bool {
	for _, item := range slice {
		if item == query {
			return true
		}
	}
	return false
}
