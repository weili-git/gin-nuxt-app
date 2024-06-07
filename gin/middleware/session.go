package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Session() gin.HandlerFunc {
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	return sessions.Sessions("my", store)
}

// session存在客户端和服务端，对于分布式系统，可能比较占用服务端的资源？
