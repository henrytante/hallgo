package webcrawler

import (
	"fmt"
	"net/http"
	"net/url"
	"golang.org/x/net/html"
	cl "github.com/fatih/color"
)
var links []string

var red = cl.New(cl.FgRed, cl.Bold)
var green = cl.New(cl.FgHiGreen, cl.Bold)

func Crawler(){
	var host string
	red.Print("Digite a url: ")
	fmt.Scanln(&host)
	url := host
	resp, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		panic(red.Sprintf("Status code diferente de 200: %v",resp.StatusCode))
	}
	doc, err := html.Parse(resp.Body)
	if err != nil{
		panic(err)
	}
	linkVisited(doc)

}
func linkVisited(doc *html.Node){
	if doc.Type == html.ElementNode || doc.Data == "a"{
		for _, atrr := range doc.Attr{
			if atrr.Key != "href"{
				continue
			}
			link, err := url.Parse(atrr.Val)
			if err != nil || link.Scheme == ""{
				continue
			}
			green.Println(link.String())
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling{
		linkVisited(c)
	}
}