package controllers

import (
	"net/http"
	"my_book/database"
	"gopkg.in/mgo.v2/bson"
	"my_book/controllers/functions"
	"html/template"
)

func ViewSearchList(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	search := r.Form["search"]

	if len(search) <= 0 || len(search[0]) <= 0 {
		functions.OutPutPageNotFound(w)
		return
	}

	dbBook := database.Book{}

	dbBookList := dbBook.FindAll(bson.M{"name": bson.M{"$regex": search[0], "$options": "$i"}})

	list := []book{}

	for _, v := range dbBookList {
		list = append(list, book{
			v.Id.Hex(),
			v.Name,
		})
	}

	t, _ := template.ParseFiles("static/searchList.html")

	t.Execute(w, list)

}