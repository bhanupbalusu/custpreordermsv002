package product_details

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

type Connection interface {
	Close()
	DB(db string) *mgo.Database
}

type conn struct {
	session *mgo.Session
}

func NewConnection(db string) Connection {
	var c conn
	var err error
	url := getURL(db)
	c.session, err = mgo.Dial(url)
	if err != nil {
		log.Panicln(err.Error())
	}
	return &c
}

func (c *conn) Close() {
	c.session.Close()
}

func (c *conn) DB(db string) *mgo.Database {
	return c.session.DB(db)
}

func getURL(db string) string {
	return fmt.Sprintf("mongodb://localhost:27017/%s", db)
}
