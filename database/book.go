package database

import (
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Book struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name"`
}

var book_collent = "books"

func (book *Book) FindId() Book {

	result := Book{}

	getMongoSession().DB(DB).C(book_collent).FindId(book.Id).One(&result)

	return result
}

func (book *Book) Insert() {

	if book.Id == "" {
		book.Id = bson.NewObjectId()
	}

	err := getMongoSession().DB(DB).C(book_collent).Insert(book)

	if err != nil {
		log.Fatal(err)
	}
}

func (book *Book) FindOne(query interface{}) Book {

	result := Book{}

	getMongoSession().DB(DB).C(book_collent).Find(query).One(&result)

	return result
}

func (book *Book) FindAll(query interface{}) []Book {

	result := []Book{}

	getMongoSession().DB(DB).C(book_collent).Find(query).All(&result)

	return result
}
