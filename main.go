package main

import (
	"easy_learning/db"
	"easy_learning/server"
	"log"
)

func main() {
	db.InitMongoConnect()

	router := server.Router()
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal("HTTP 服务开启失败：", err)
	}
}
