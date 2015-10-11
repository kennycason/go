package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
)

type Link struct {
    Title string
    Url string
}

func parseLinks(doc *goquery.Document) [] Link {
    linkNodes := doc.Find(".athing")
    links := make([]Link, len(linkNodes.Nodes))

    linkNodes.Each(func(i int, s *goquery.Selection) {
        title := s.Find(".title").Text()
        url, _ := s.Find(".title a").Attr("href")

        links[i] = Link{Title: title, Url: url}
    })
    return links
}

func main() {
    doc, _ := goquery.NewDocument("https://news.ycombinator.com/")
    
    links := parseLinks(doc)
    fmt.Println(links)
}