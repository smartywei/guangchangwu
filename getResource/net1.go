package getResource

import (
	"my_book/request"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"regexp"
)

type Book struct {
	Name string
	Href string
}

type Catlog struct {
	Order int
	Name  string
	Href  string
}

type Content struct {
	Content string
}

func GetSearch(search_word string) []Book {

	search_href := "http://zhannei.baidu.com/cse/search?q=" + search_word + "&s=13049992925692302651&entry=1"

	request := &bookrequest.Request{
		search_href,
		nil,
		nil,
	}

	response := request.Get()

	//body := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
	body := response.Body

	dom, err := goquery.NewDocumentFromReader(body)

	response.Body.Close()

	if err != nil {
		panic(err)
	}

	bookList := []Book{}

	dom.Find(".result-item h3 a").Each(func(i int, selection *goquery.Selection) {

		title := strings.TrimSpace(selection.Text())
		href, _ := selection.Attr("href")
		href = strings.TrimSpace(href)

		bookList = append(bookList, Book{
			title,
			href,
		})
	})

	return bookList
}

func GetCatlogs(href string) []Catlog {

	request := &bookrequest.Request{
		href,
		nil,
		nil,
	}

	response := request.Get()

	body := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())

	dom, err := goquery.NewDocumentFromReader(body)

	response.Body.Close()

	if err != nil {
		panic(err)
	}

	catLogList := []Catlog{}

	dom.Find(".chapterlist li a").Each(func(i int, selection *goquery.Selection) {
		title := strings.TrimSpace(selection.Text())
		link, _ := selection.Attr("href")
		link = strings.TrimSpace(link)

		catLogList = append(catLogList, Catlog{
			i,
			title,
			href + link,
		})
	})

	return catLogList

}

func GetContents(href string) Content {

	request := &bookrequest.Request{
		href,
		nil,
		nil,
	}

	response := request.Get()

	body := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())

	dom, err := goquery.NewDocumentFromReader(body)

	response.Body.Close()

	if err != nil {
		panic(err)
	}

	content, _ := dom.Find("#htmlContent").Html()

	r, _ := regexp.Compile("\n                        武林中文网 WWW.50ZW.LA，最快更新(.*?)最新章节！<br/><br/>")

	s := r.FindString(content)

	return Content{
		strings.Replace(content, s, "", 1),
	}

}