package database

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"time"
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

	err := getMongoSession().DB(DB).C(catlog_collent).FindId(catlog.Id).One(&result)

	if err != nil {
		fmt.Println("根据ID查询目录失败，正在重试...")
		time.Sleep(time.Second*5)
		return catlog.FindId()
	}

	return result
}

func (catlog *Catlog) Insert() bson.ObjectId{

	if catlog.Id == "" {
		catlog.Id = bson.NewObjectId()
	}

	err := getMongoSession().DB(DB).C(catlog_collent).Insert(catlog)

	if err != nil {
		fmt.Println("插入目录失败，正在重试...")
		time.Sleep(time.Second*5)
		return catlog.Insert()
	}

	return catlog.Id
}

func (catlog *Catlog) FindOne(query interface{}) Catlog {

	result := Catlog{}

	err := getMongoSession().DB(DB).C(catlog_collent).Find(query).One(&result)

	if err != nil {
		fmt.Println("查询单个目录失败，正在重试...")
		time.Sleep(time.Second*5)
		return catlog.FindOne(query)
	}

	return result
}

func (catlog *Catlog) FindAll(query interface{}) []Catlog {

	result := []Catlog{}

	err := getMongoSession().DB(DB).C(catlog_collent).Find(query).All(&result)

	if err != nil {
		fmt.Println("查询目录列表失败，正在重试...")
		time.Sleep(time.Second*5)
		return catlog.FindAll(query)
	}

	return result
}
