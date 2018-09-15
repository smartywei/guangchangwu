package main

import (
	"my_book/getResource"
	"gopkg.in/mgo.v2/bson"
	"my_book/database"
	"fmt"
	"time"
)

func main() {

	bookList := getResource.GetSearch("天道图书馆")

	if len(bookList) <= 0 {
		panic("没有搜索内容")
	}

	book := bookList[0]

	dbBook := database.Book{
		Name: book.Name,
	}

	book_id := dbBook.Insert()

	fmt.Println("书籍插入成功： ID ------》", book_id)

	time.Sleep(time.Second * 3)

	catlogList := getResource.GetCatlogs(book.Href)


	for _, v := range catlogList {

		saveResourceToMongoDB(book_id, v, getResource.Content{
			Content: getResource.GetContents(v.Href).Content,
		})

		time.Sleep(time.Second * 3)
	}

	database.Close()
}

func saveResourceToMongoDB(book_id bson.ObjectId, catlog getResource.Catlog, content getResource.Content) {
	//1.插入目录表
	dbCatLog := database.Catlog{
		Catlog:  catlog.Name,
		Order:   catlog.Order,
		Book_id: book_id,
	}

	cat_id := dbCatLog.Insert()

	fmt.Println("目录插入成功： ID ------》", cat_id ,"OrderId-------》",catlog.Order)

	dbContent := database.Content{
		Order:        catlog.Order,
		Cat_log_id:   cat_id,
		Cat_log_name: catlog.Name,
		Book_id:      book_id,
		Content:      content.Content,
	}

	content_id := dbContent.Insert()

	fmt.Println("内容插入成功： ID ------》", content_id,"OrderId-------》",catlog.Order)
}
