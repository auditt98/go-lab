package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

var urlMap = map[string]struct{}{}

func validate(uri string) {
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		panic("Invalid URL")
	}
}

func crawl(uri string) {
	resp, err := http.Get(uri)
	if err != nil {
		panic("Error fetching URL: " + err.Error())
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic("Can't parse HTML")
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			var linkUrl string
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					if strings.HasPrefix(attr.Val, "/") {
						linkUrl = uri + attr.Val
					} else {
						linkUrl = attr.Val
					}
					if strings.HasSuffix(linkUrl, "/") {
						linkUrl = linkUrl[:len(linkUrl)-1]
					}

					break
				}
			}
			if _, ok := urlMap[linkUrl]; ok {
				return
			} else {
				urlMap[linkUrl] = struct{}{}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func main() {
	var inp string
	fmt.Print("Enter an URL: ")
	_, err := fmt.Scanln(&inp)
	if err != nil {
		panic("Error reading input")
	}
	validate(inp)
	crawl(inp)
	for k := range urlMap {
		fmt.Println(k)
	}
}
