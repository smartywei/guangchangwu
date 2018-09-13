package controllers

import (
	"net/http"
	"my_book/database"
	"html/template"
	"my_book/controllers/functions"
	"gopkg.in/mgo.v2/bson"
)

type content struct {
	Cat_log_name template.HTML `json:"cat_log_name"`
	Content      template.HTML `json:"content"`
	Book_id      string        `json:"book_id"`
	Order        int           `json:"order"`
	PreHref      string        `json:"pre_href"`
	NextHref     string        `json:"next_href"`
}

func ViewContentList(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	cat_id := r.Form["cat_id"]

	if len(cat_id) <= 0 || len(cat_id[0]) <= 0 {
		functions.OutPutPageNotFound(w)
		return
	}

	dbContent := database.Content{}

	dbContentInfo := dbContent.FindOne(bson.M{"cat_log_id": bson.ObjectIdHex(cat_id[0])})


	var preHref string
	var nextHref string

	if dbContentInfo.Order > 0 {
		dbPreContentInfo := dbContent.FindOne(bson.M{"book_id": dbContentInfo.Book_id,"order":dbContent.Order-1})
		preHref = "/content?cat_id="+dbPreContentInfo.Cat_log_id.Hex()
	}else{
		preHref = "#"
	}


	dbNextContentInfo := dbContent.FindOne(bson.M{"book_id": dbContentInfo.Book_id,"order":dbContent.Order+1})

	nextHref = "/content?cat_id="+dbNextContentInfo.Cat_log_id.Hex()

	t, _ := template.ParseFiles("static/content.html")

	content_res := content{
		template.HTML(dbContentInfo.Cat_log_name),
		template.HTML(dbContentInfo.Content),
		dbContentInfo.Book_id.Hex(),
		dbContentInfo.Order,
		preHref,
		nextHref,
	}

	t.Execute(w, content_res)
}
