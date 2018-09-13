package controllers

import (
	"my_book/database"
	"net/http"
	"my_book/controllers/functions"
	"html/template"
)

type book struct {
	Id   string
	Name string
}

func ViewBookList(w http.ResponseWriter, r *http.Request)  {

	r.ParseForm()

	if r.URL.Path != "/" {
		functions.OutPutPageNotFound(w)
		return
	}

	dbBook := database.Book{}

	dbBookList := dbBook.FindAll(nil)

	list := []book{}

	for _, v := range dbBookList {
		list = append(list, book{
			v.Id.Hex(),
			v.Name,
		})
	}

	t, _ := template.ParseFiles("static/bookList.html")

	t.Execute(w, list)
}
