package controllers

import (
	"net/http"
	"my_book/controllers/functions"
	"my_book/database"
	"html/template"
	"gopkg.in/mgo.v2/bson"
)

type catLog struct {
	Id    string
	Name  string
	Order int
}

type catLogList struct{
	CatLogList []catLog
	BookInfo   database.Book
}


func ViewCatLogList(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	book_id := r.Form["book_id"]

	if len(book_id) <= 0 || len(book_id[0]) <= 0 {
		functions.OutPutPageNotFound(w)
		return
	}

	dbBook := database.Book{
		Id: bson.ObjectIdHex(book_id[0]),
	}

	dbBookInfo := dbBook.FindId()

	dbCatLog := database.Catlog{}

	dbCatLogList := dbCatLog.FindAll(bson.M{"book_id": bson.ObjectIdHex(book_id[0])})

	list := []catLog{}

	for _, v := range dbCatLogList {
		list = append(list, catLog{
			v.Id.Hex(),
			v.Catlog,
			v.Order,
		})
	}

	t, _ := template.ParseFiles("static/catLogList.html")

	catLogInfo := catLogList{
		list,
		dbBookInfo,
	}

	t.Execute(w, catLogInfo)

}
