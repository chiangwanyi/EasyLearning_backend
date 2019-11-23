package main

import (
	"easy_learning/config"
	"easy_learning/db"
	"easy_learning/server"
	"log"
)

func main() {
	// 初始化数据库连接
	db.InitMongoConnect()

	// 注册路由
	router := server.Router()

	// 开启 HTTP 服务
	if err := router.Run(config.HttpAddress); err != nil {
		log.Fatal("HTTP 服务开启失败：", err)
	}
}
