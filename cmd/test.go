package main

import (
	"book/getResource"
	"fmt"
)

//var pcHeader = map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"}

func main() {

	//request := &bookrequest.Request{
	//	"http://zhannei.baidu.com/cse/search?q=天道图书馆&s=13049992925692302651&entry=1",
	//	nil,
	//	nil,
	//}
	//
	//response := request.Get()
	//
	////body := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
	//body := response.Body
	//
	//dom, err := goquery.NewDocumentFromReader(body)
	//
	//response.Body.Close()
	//
	//if err != nil{
	//	panic(err)
	//}
	//
	//dom.Find(".result-item h3 a").Each(func(i int, selection *goquery.Selection) {
	//	fmt.Println(selection.Text())
	//	fmt.Print(selection.Attr("href"))
	//})

	//str,_ := dom.Html()
	//
	//fmt.Println(str)

	//bookList := getResource.GetSearch("天道图书馆")
	//for _,v := range bookList{
	//	fmt.Println(v.Name)
	//	fmt.Println(v.Href)
	//}

	//catList := getResource.GetCatlogs("http://www.50zw.la/book_80860/")
	//for _,v := range catList{
	//	fmt.Println(v.Order,"  ",v.Name,"   ",v.Href)
	//}

	content := getResource.GetContents("http://www.50zw.la/book_80860/25066376.html")

	fmt.Println(content.Content)
}