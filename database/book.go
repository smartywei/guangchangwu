package database

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"time"
)

type Book struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name"`
}

var book_collent = "books"

func (book *Book) FindId() Book {

	result := Book{}

	err := getMongoSession().DB(DB).C(book_collent).FindId(book.Id).One(&result)

	if err != nil {
		fmt.Println("根据ID查询书籍失败，正在重试...")
		time.Sleep(time.Second * 5)
		return book.FindId()
	}

	return result
}

func (book *Book) Insert() bson.ObjectId {

	if book.Id == "" {
		book.Id = bson.NewObjectId()
	}

	err := getMongoSession().DB(DB).C(book_collent).Insert(book)

	if err != nil {
		fmt.Println("插入书籍失败，正在重试...")
		time.Sleep(time.Second * 5)
		return book.Insert()
	}

	return book.Id
}

func (book *Book) FindOne(query interface{}) Book {

	result := Book{}

	err := getMongoSession().DB(DB).C(book_collent).Find(query).One(&result)

	if err != nil {
		fmt.Println("查询单个书籍失败，正在重试...")
		time.Sleep(time.Second * 5)
		return book.FindOne(query)
	}

	return result
}

func (book *Book) FindAll(query interface{}) []Book {

	result := []Book{}

	err := getMongoSession().DB(DB).C(book_collent).Find(query).All(&result)

	if err != nil {
		fmt.Println("查询书籍列表失败，正在重试...")
		time.Sleep(time.Second * 5)
		return book.FindAll(query)
	}

	return result
}
