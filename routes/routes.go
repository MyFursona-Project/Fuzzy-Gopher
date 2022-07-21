package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func SetRoutes(router *gin.Engine, pg_conn *gorm.DB, redisVerify, redisSession *redis.Client) {
	// anonymous fursona creation
	router.POST("/anonymous", createAnonymous()) // create an anonymous session
	router.POST("/anonymous/create", getHealth)  // create a fursona item

	// normal fursona interaction
	router.POST("/fursona", getHealth)            // create a fursona
	router.GET("/fursona/trending", getHealth)    // get trending fursonas
	router.POST("/fursona/trending", getHealth)   // get trending fursonas, signed in
	router.GET("/fursona/recent", getHealth)      // get recent fursonas
	router.POST("/fursona/recent", getHealth)     // get recent fursonas, signed in
	router.GET("/fursona/:id", getHealth)         // get a specific fursona
	router.POST("/fusona/:id", getHealth)         // get specific fursona while signed in
	router.PATCH("/fursona/:id", getHealth)       // update a fursona
	router.POST("/fursona/:id/remove", getHealth) // remove a sona
	router.POST("/fursona/:id/like", getHealth)   // like a specific fursona
	router.POST("/fursona/:id/unlike", getHealth) // remove a like from a sona

	// content reports
	router.POST("/fursona/:id/report", getHealth) // create a content report for a fursona
	router.POST("/profile/report", getHealth)     // report an account

	// profile interaction
	router.POST("/profile", getHealth)       // get the currently signed in
	router.PATCH("/profile", getHealth)      // update the signed in profile
	router.GET("/profile/:name", getHealth)  // fetch a profile
	router.POST("/profile/:name", getHealth) // fetch a profile signed in

	router.Static("/swagger", "swagger/")
	// user redirects
	router.GET("/", redirect("/swagger"))
	router.GET("/v1", redirect("/swagger"))
	// static responses
	router.GET("/health", getHealth)
}
