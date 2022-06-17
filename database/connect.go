package database

import (
	"fmt"
	"os"

	"github.com/MyFursona-Project/Fuzzy-Gopher/tools"
	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectSQL() *gorm.DB {
	// build connection URI
	uri := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASS"),
		os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_SSLMODE"),
		os.Getenv("POSTGRES_TIMEZONE"))

	// connect to the DB
	// TODO: does 'warn' loglevel contain sensitive data?
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{Logger: logger.Default.LogMode(logger.Warn)})
	if err != nil {
		tools.ErrorFatal("Libpuroto", err)
	}
	return db
}
func AutoMigrateSQL(pg_conn *gorm.DB) {
	if err := pg_conn.AutoMigrate(&User{}, &Verify{}, &Profile{}, &Fursona{}, &FursonaMedia{}, &Like{}); err != nil {
		tools.ErrorPanic(err)
	}
}

func ConnectRedis(dbNumber int) *redis.Client {
	// return redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: dbNumber})
	return redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_HOST"), Password: os.Getenv("REDIS_PASS"), DB: dbNumber})
}
