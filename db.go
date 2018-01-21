package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DBURL      = "localhost:27017"
	DBNAME     = "test"
	COLLECTION = "posts"
)

type Post struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Date  int64         `bson:"date"`
}

func open() (*mgo.Session, error) {
	session, err := mgo.Dial(DBURL)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func getPosts(after int) ([]Post, error) {
	s, err := open()
	if err != nil {
		return nil, err
	}
	defer s.Close()

	c := s.DB(DBNAME).C(COLLECTION)
	//db.posts.find({"date": { "$gte": 0 }}).sort({"date": 1}).limit()
	var result []Post
	err = c.Find(bson.M{"date": bson.M{"$gte": after}}).Sort("-date").Limit(10).All(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
