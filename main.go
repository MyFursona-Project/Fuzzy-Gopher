package main

import (
	"github.com/MyFursona-Project/Fuzzy-Gopher/database"
	"github.com/MyFursona-Project/Fuzzy-Gopher/routes"
	"github.com/MyFursona-Project/Fuzzy-Gopher/tools"
	"github.com/gin-gonic/gin"
)

func main() {
	// connect to the PostgreSQL
	pg_conn := database.ConnectSQL()
	if pg_conn.Error != nil {
		tools.ErrorPanic(pg_conn.Error)
	}
	// Connect to Redis
	redisVerify := database.ConnectRedis(1)
	redisSession := database.ConnectRedis(2)

	// check if redis can be reached
	if err := redisVerify.Ping().Err(); err != nil {
		tools.ErrorPanic(err)
	} else if err := redisSession.Ping().Err(); err != nil {
		tools.ErrorPanic(err)
	}

	// migrate all tables
	database.AutoMigrateSQL(pg_conn)

	// create router
	router := gin.Default()

	// TODO: set trusted proxies for prod

	// set routes
	routes.SetRoutes(router, pg_conn, redisVerify, redisSession)

	// start
	if err := router.Run("0.0.0.0:3926"); err != nil {
		tools.ErrorPanic(err)
	}

	// clean up
	redisVerify.Close()
	redisSession.Close()
}
