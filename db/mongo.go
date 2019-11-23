package db

import (
	"easy_learning/config"
	"github.com/globalsign/mgo"
	"log"
)

// Mongo MongoDB 连接
var MongoSession *mgo.Session

//InitMongoConnect 初始化 MongoDB 数据库连接
func InitMongoConnect() {
	m, err := mgo.Dial(config.MongoDBAddress)
	if err != nil {
		log.Fatal("初始化 MongoDB 数据库失败：", err)
	}
	MongoSession = m
}