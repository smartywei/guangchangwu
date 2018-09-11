package main

import (
	"net/http"
	"log"
	"html/template"
	"my_book/database"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type BookList struct {
	Id   string
	Name string
}

type CatLogList struct {
	Id    string
	Name  string
	Order int
}

type catLogInfo struct {
	CatLogList []CatLogList
	BookInfo   database.Book
}

type Content struct {
	Cat_log_name template.HTML `json:"cat_log_name"`
	Content      template.HTML `json:"content"`
	Book_id string `json:"book_id"`
	Order int `json:"order"`
}

func bookList(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	if r.URL.Path != "/" {
		fmt.Fprintln(w, "404 page not fout")
		return
	}

	book := database.Book{}

	bookList := book.FindAll(nil)

	list := []BookList{}

	for _, v := range bookList {

		list = append(list, BookList{
			v.Id.Hex(),
			v.Name,
		})
	}

	t, _ := template.ParseFiles("static/bookList.html")

	t.Execute(w, list)
}

func catList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	book_id := r.Form["book_id"]
	if len(book_id) <= 0 || len(book_id[0]) <= 0 {
		fmt.Fprint(w, "404 page not fout")
	} else {

		book := database.Book{
			Id: bson.ObjectIdHex(book_id[0]),
		}

		book_info := book.FindId()

		catLog := database.Catlog{}

		catLogList := catLog.FindAll(bson.M{"book_id": bson.ObjectIdHex(book_id[0])})

		list := []CatLogList{}

		for _, v := range catLogList {
			list = append(list, CatLogList{
				v.Id.Hex(),
				v.Catlog,
				v.Order,
			})
		}

		t, _ := template.ParseFiles("static/catLogList.html")

		catLogInfo := catLogInfo{
			list,
			book_info,
		}

		t.Execute(w, catLogInfo)

	}

}

func content(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cat_id := r.Form["cat_id"]
	if len(cat_id) <= 0 || len(cat_id[0]) <= 0 {
		fmt.Fprint(w, "404 page not fout")
	} else {
		content := database.Content{}

		content_info := content.FindOne(bson.M{"cat_log_id": bson.ObjectIdHex(cat_id[0])})

		t, _ := template.ParseFiles("static/content.html")

		content_res := Content{
			template.HTML(content_info.Cat_log_name),
			template.HTML(content_info.Content),
			content_info.Book_id.Hex(),
			content_info.Order,
		}

		t.Execute(w, content_res)
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	search := r.Form["search"]
	if len(search) <= 0 || len(search[0]) <= 0 {
		fmt.Fprint(w, "404 page not fout")
		return
	}

	book := database.Book{}

	bookList := book.FindAll(bson.M{"name": bson.M{"$regex": search[0], "$options": "$i"}})

	list := []BookList{}

	for _, v := range bookList {

		list = append(list, BookList{
			v.Id.Hex(),
			v.Name,
		})
	}

	t, _ := template.ParseFiles("static/bookList.html")

	t.Execute(w, list)

}

func main() {
	http.HandleFunc("/", bookList)       // 设置访问的路由
	http.HandleFunc("/catlist", catList) // 设置访问的路由
	http.HandleFunc("/content", content) // 设置访问的路由
	http.HandleFunc("/search", Search)   // 设置访问的路由

	err := http.ListenAndServe(":80", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
