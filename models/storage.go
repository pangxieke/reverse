package models

import (
	"github.com/globalsign/mgo"
)

var Mgo *StorageMgo

type StorageMgo struct {
	db      string
	session *mgo.Session
}

func NewStorage(db string, session *mgo.Session) *StorageMgo {
	return &StorageMgo{
		db:      db,
		session: session,
	}
}

func MgoAdd(collection string, data interface{}) (err error) {
	ss := Mgo.session.Copy()
	defer ss.Close()

	err = ss.DB(Mgo.db).C(collection).Insert(data)
	return
}
