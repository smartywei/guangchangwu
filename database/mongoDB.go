package database

import (
	"gopkg.in/mgo.v2"
)

var Session *mgo.Session

var DB = "book"

func getMongoSession() *mgo.Session {

	if Session == nil {

		sess, err := mgo.Dial("mongodb://root:123456@127.0.0.1:27017")

		Session = sess

		if err != nil {
			panic(err)
		}

		Session.SetMode(mgo.Monotonic, true)
	}

	return Session
}

func Close() {
	Session.Close()
}