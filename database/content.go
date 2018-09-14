package database

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"time"
)

type Content struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`
	Order        int           `json:"order"`
	Cat_log_id   bson.ObjectId `json:"cat_log_id"`
	Cat_log_name string        `json:"cat_log_name"`
	Book_id      bson.ObjectId `json:"book_id"`
	Content      string        `json:"content"`
}

var content_collent = "contents"

func (content *Content) FindId() Content {

	result := Content{}

	err := getMongoSession().DB(DB).C(content_collent).FindId(content.Id).One(&result)

	if err != nil {
		fmt.Println("根据ID查询内容失败，正在重试...")
		time.Sleep(time.Second * 5)
		return content.FindId()
	}

	return result
}

func (content *Content) Insert() bson.ObjectId {

	if content.Id == "" {
		content.Id = bson.NewObjectId()
	}

	err := getMongoSession().DB(DB).C(content_collent).Insert(content)

	if err != nil {
		fmt.Println("插入内容失败，正在重试...")
		time.Sleep(time.Second * 5)
		return content.Insert()
	}

	return content.Id
}

func (content *Content) FindOne(query interface{}) Content {

	result := Content{}

	err := getMongoSession().DB(DB).C(content_collent).Find(query).One(&result)

	if err != nil {
		fmt.Println("查询单个目录失败，正在重试...")
		time.Sleep(time.Second * 5)
		return content.FindOne(query)
	}

	return result
}

func (content *Content) FindAll(query interface{}) []Content {

	result := []Content{}

	err := getMongoSession().DB(DB).C(content_collent).Find(query).All(&result)

	if err != nil {
		fmt.Println("查询目录列表失败，正在重试...")
		time.Sleep(time.Second*5)
		return content.FindAll(query)
	}

	return result
}
