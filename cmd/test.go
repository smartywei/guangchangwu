package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"book/request"
)

//var pcHeader = map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"}

func main() {

	request := &bookrequest.Request{
		"https://studygolang.com/",
		nil,
		nil,
	}

	response := request.Get()

	dom, err := goquery.NewDocumentFromReader(response.Body)

	response.Body.Close()

	if err != nil{
		panic(err)
	}

	fmt.Println(dom.Html())
}
