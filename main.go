package main

import (
	"fmt"
	"os"

	"github.com/gaku3601/golang-and-firebaseauth/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	gdb := initDB()
	r := gin.Default()
	r.Use(apiMiddleware(gdb))
	profiles := r.Group("/profiles")
	{
		controllers.GetProfiles(profiles)
		controllers.PostProfiles(profiles)
	}
	r.Run()
}

func initDB() *gorm.DB {
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DB")
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName))
	if err != nil {
		panic(err)
	}
	return db
}

func apiMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConn", db)
		c.Next()
	}
}
