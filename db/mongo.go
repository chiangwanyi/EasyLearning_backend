package db

import (
	"github.com/globalsign/mgo"
	"log"
)

// Mongo MongoDB 连接
var MongoSession *mgo.Session

//InitMongoConnect 初始化 MongoDB 数据库连接
func InitMongoConnect() {
	m, err := mgo.Dial("localhost:27017/el")
	if err != nil {
		log.Fatal("初始化 MongoDB 数据库失败：", err)
	}
	MongoSession = m
}