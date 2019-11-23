package middleware

import (
	"easy_learning/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/mongo"
	"github.com/gin-gonic/gin"
)

// Sessions 初始化 Session
func Sessions() gin.HandlerFunc {
	session := db.MongoSession.Copy()

	client := session.DB("").C("session")
	store := mongo.NewStore(client, 3600, true, []byte("secret"))
	store.Options(sessions.Options{
		HttpOnly: true,
	})
	return sessions.Sessions("my_session", store)
}
