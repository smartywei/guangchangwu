package database

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"time"
)

var Session *mgo.Session

var DB = "book"

func getMongoSession() *mgo.Session {

	if Session == nil {

		sess, err := mgo.Dial("mongodb://smartywei:123456@118.24.43.196:27017")

		Session = sess

		if err != nil {
			fmt.Println("数据库连接失败，正在重新尝试连接数据库...")
			time.Sleep(time.Second * 10)
			return getMongoSession()
		}

		Session.SetMode(mgo.Monotonic, true)
	}

	return Session
}

func Close() {
	Session.Close()
}