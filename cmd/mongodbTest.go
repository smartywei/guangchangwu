package main

import (
	"book/database"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	book := database.Book{
		Name: "天道图书馆",
	}

	book.Insert()

	res1 := book.FindOne(bson.M{"name": "天道图书馆"})

	fmt.Println(res1)

	catlog := database.Catlog{
		Catlog:  "第一章",
		Order:   1,
		Book_id: res1.Id,
	}

	catlog.Insert()

	res2 := catlog.FindOne(bson.M{"catlog": "第一章"})

	fmt.Println(res2)

	content := database.Content{
		Cat_log_id: res2.Id,
		Content:    "赵日天无敌",
	}

	content.Insert()

	res3 := content.FindOne(bson.M{"cat_log_id": res2.Id})

	fmt.Println(res3)

	database.Close()
}
