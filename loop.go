// package main

// import "fmt"

// func main() {
// 	a := 3
// 	for x:=5; a <25; x+=3{
// 		fmt.Println("Do stuff",x)
// 		a+=4
// 	}
// }

// ________________________________________________

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)


type Sitemapindex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	var s Sitemapindex
	xml.Unmarshal(bytes, &s)
	// fmt.Println(s.Locations)

	for _, Location := range s.Locations {
		fmt.Printf("%s\n", Location)
	}
}