package main

import (
	"my_book/getResource"
	"my_book/database"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	content := getResource.GetContents("http://m.50zw.la/book_80860/25104722.html");

		dbContent := database.Content{
			Order:        1,
			Cat_log_id:   bson.NewObjectId(),
			Cat_log_name: "aaa",
			Book_id:     bson.NewObjectId(),
			Content:      content.Content,
		}

		dbContent.Insert()
}

//func saveResourceToMongoDB(book_id bson.ObjectId, catlog getResource.Catlog, content getResource.Content) {
//	//1.插入目录表
//	dbCatLog := database.Catlog{
//		Catlog:  catlog.Name,
//		Order:   catlog.Order,
//		Book_id: book_id,
//	}
//
//	cat_id := dbCatLog.Insert()
//
//	fmt.Println("目录插入成功： ID ------》", cat_id ,"OrderId-------》",catlog.Order)
//
//	dbContent := database.Content{
//		Order:        catlog.Order,
//		Cat_log_id:   cat_id,
//		Cat_log_name: catlog.Name,
//		Book_id:      book_id,
//		Content:      content.Content,
//	}
//
//	content_id := dbContent.Insert()
//
//	fmt.Println("内容插入成功： ID ------》", content_id,"OrderId-------》",catlog.Order)
//}
