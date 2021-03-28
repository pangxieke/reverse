package models

import (
	"fmt"
	"github.com/globalsign/mgo"
	"reverse/config"
	"reverse/log"
)

func Init() (err error) {
	if err = initMgo(); err != nil {
		return err
	}
	return
}

func initMgo() (err error) {
	url := config.Mgo.Urls
	db := config.Mgo.DB
	if url == "" || db == "" {
		panic(fmt.Sprintf("mongodb init error, url=%s, db=%s", url, db))
	}
	log.Info("initialize storage, url = %s", url)

	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	Mgo = NewStorage(db, session)
	return
}
