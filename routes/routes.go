package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func SetRoutes(router *gin.Engine, pg_conn *gorm.DB, redisVerify, redisSession *redis.Client) {
	router.Static("/swagger", "swagger/")
	// user redirects
	router.GET("/", redirect("/swagger"))
	router.GET("/v1", redirect("/swagger"))
	// static responses
	router.GET("/health", getHealth)
}
