package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	// Titles []string `xml:"url>news>title"`
	// Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// type NewsMap struct {
// 	Keyword  string
// 	Location string
// }

func main() {
	var s SitemapIndex
	// var n News
	// news_map := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	// fmt.Println(string_body)
	// resp.Body.Close()

	xml.Unmarshal(bytes, &s)

	// fmt.Println(s.Locations)
	for _, Location := range s.Locations {
		fmt.Println(Location)
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(bytes)
		// xml.Unmarshal(bytes, &n)
		// fmt.Println(n)
		// for idx, _ := range n.Titles {
		// 	news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		// }
	}

	// for idx, data := range news_map {
	// 	fmt.Println("\n\n\n", idx)
	// 	fmt.Println("\n", data.Keyword)
	// 	fmt.Println("\n", data.Location)
	// }
}
