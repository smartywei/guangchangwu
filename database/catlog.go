package database

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"fmt"
)

type Catlog struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Book_id bson.ObjectId `json:"book_id"`
	Order int `json:"order"`
	Catlog string        `json:"catlog"`
}

var catlog_collent = "catalogs"

func (catlog *Catlog) FindId() Catlog {

	result := Catlog{}

	getMongoSession().DB(DB).C(catlog_collent).FindId(catlog.Id).One(&result)

	return result
}

func (catlog *Catlog) Insert() bson.ObjectId{

	if catlog.Id == "" {
		catlog.Id = bson.NewObjectId()
	}

	err := getMongoSession().DB(DB).C(catlog_collent).Insert(catlog)

	if err != nil {
		fmt.Println(2)
		log.Fatal(err)
	}

	return catlog.Id
}

func (catlog *Catlog) FindOne(query interface{}) Catlog {

	result := Catlog{}

	getMongoSession().DB(DB).C(catlog_collent).Find(query).One(&result)

	return result
}

func (catlog *Catlog) FindAll(query interface{}) []Catlog {

	result := []Catlog{}

	getMongoSession().DB(DB).C(catlog_collent).Find(query).All(&result)

	return result
}
