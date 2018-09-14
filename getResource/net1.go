package getResource

import (
	"my_book/request"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"regexp"
	"strconv"
	"fmt"
	"time"
)

var mBaseHref = "http://m.50zw.la"

var mCatLogHref = "http://m.50zw.la/chapters_"

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

		r,_ := regexp.Compile("/book_([\\d]+)/")

		bookId := string([]byte(r.FindString(href))[6:])

		bookList = append(bookList, Book{
			title,
			mCatLogHref + bookId,
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

	var index_page = 1

	response := request.Get()

	body := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())

	dom, err := goquery.NewDocumentFromReader(body)

	response.Body.Close()

	pageCountR,_ := regexp.Compile("/([\\d]+)")

	pageCountStr :=    dom.Find(".page-book-turn").Text()

	pageCountIndex := pageCountR.FindStringIndex(pageCountStr)

	pageCount,_ :=  strconv.Atoi(string([]byte(pageCountStr)[pageCountIndex[0]+1:pageCountIndex[1]]))

	fmt.Println("当前书籍一共：",pageCount,"页目录")

	if err != nil {
		panic(err)
	}

	catLogList := []Catlog{}

	catTitleKyeList := map[string]int{}

	var i = 0

	for  index_page <=  pageCount {

		dom.Find(".last9 li a").Each(func(index int, selection *goquery.Selection) {

			if index == 0{
				return
			}

			title := strings.TrimSpace(selection.Text())
			link, _ := selection.Attr("href")
			link = strings.TrimSpace(link)

			r, _ := regexp.Compile("第(.*?)章(.*?)")

			if !r.MatchString(title) {
				return
			}

			r2, _ := regexp.Compile("第([\\d]+)章(.*?)")

			if r2.MatchString(title) {
				return
			}

			r3,_ := regexp.Compile("第(.*?)章")

			tagString := r3.FindString(title)

			if catTitleKyeList[tagString] != 0 {
				return
			}

			catLogList = append(catLogList, Catlog{
				i,
				title,
				mBaseHref+link,
			})

			catTitleKyeList[tagString] = i

			i++
		})

		fmt.Println("第"+strconv.Itoa(index_page)+"页目录抓取完毕....")

		time.Sleep(time.Second * 3)

		index_page++

		request.Url = href + strconv.Itoa(index_page)

		response := request.Get()

		body = transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())

		dom,err = goquery.NewDocumentFromReader(body)

		if err != nil {
			panic(err)
		}

		response.Body.Close()


	}

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

	content, _ := dom.Find("#nr1").Html()
	nextText := dom.Find("#pb_next").Text()
	nextHref,_ := dom.Find("#pb_next").Attr("href")

	content = strings.Replace(content, "<div class=\"kongwen\"></div>", "", 1)
	content = strings.Replace(content, "<div class=\"middlead\"><script type=\"text/javascript\">_Middle();</script></div>", "", 1)
	content = strings.Replace(content, " --&gt;&gt;<br/><center class=\"red\">本章未完，点击下一页继续阅读</center>", "", 1)
	content = strings.Replace(content, " ", "&nbsp", -1)
	content = strings.Replace(content, " ", "", -1)

	fmt.Println( "小章节抓取完毕..")

	for nextText == "下一页" {

		time.Sleep(time.Second * 3)

		request.Url = mBaseHref + nextHref

		response := request.Get()

		body := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())

		dom, err = goquery.NewDocumentFromReader(body)

		response.Body.Close()

		if err != nil {
			panic(err)
		}

		nextContent, _ := dom.Find("#nr1").Html()
		nextText = dom.Find("#pb_next").Text()
		nextHref,_ = dom.Find("#pb_next").Attr("href")

		nextContent = strings.Replace(nextContent, "<div class=\"kongwen\"></div>", "", 1)
		nextContent = strings.Replace(nextContent, "<div class=\"middlead\"><script type=\"text/javascript\">_Middle();</script></div>", "", 1)
		nextContent = strings.Replace(nextContent, " --&gt;&gt;<br/><center class=\"red\">本章未完，点击下一页继续阅读</center>", "", 1)
		nextContent = strings.Replace(nextContent, " ", "&nbsp", -1)
		nextContent = strings.Replace(nextContent, " ", "", -1)
		//nextContent = strings.Replace(nextContent, "\n", "", -1)


		content = content+nextContent

		fmt.Println( "小章节抓取完毕..")
	}

	content = strings.Replace(content, "&amp;", "&", -1)
	content = strings.Replace(content, "\n", "", -1)
	content = strings.TrimSpace(content)

	//r, _ := regexp.Compile("\n                        武林中文网 WWW.50ZW.LA，最快更新(.*?)最新章节！<br/><br/>")
	//
	//s := r.FindString(content)

	return Content{
		content,
	}

}
