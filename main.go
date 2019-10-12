package main

import (
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
	db, err := gorm.Open("postgres", "host=localhost user=user password=password dbname=db sslmode=disable")
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
