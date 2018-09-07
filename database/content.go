package database

import (
"gopkg.in/mgo.v2/bson"
"log"
	"fmt"
)

type Content struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Cat_log_id bson.ObjectId `json:"cat_log_id"`
	Content string        `json:"content"`
}

var content_collent = "contents"

func (content *Content) FindId() Content {

	result := Content{}

	getMongoSession().DB(DB).C(content_collent).FindId(content.Id).One(&result)

	return result
}

func (content *Content) Insert() {

	if content.Id == "" {
		content.Id = bson.NewObjectId()
	}

	err := getMongoSession().DB(DB).C(content_collent).Insert(content)

	if err != nil {
		fmt.Println(3)
		log.Fatal(err)
	}
}

func (content *Content) FindOne(query interface{}) Content {

	result := Content{}

	getMongoSession().DB(DB).C(content_collent).Find(query).One(&result)

	return result
}

func (content *Content) FindAll(query interface{}) []Content {

	result := []Content{}

	getMongoSession().DB(DB).C(content_collent).Find(query).All(&result)

	return result
}