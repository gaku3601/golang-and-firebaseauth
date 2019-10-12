package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Profile struct
type Profile struct {
	Name string
}

//CreateProfile create
func CreateProfile(c *gin.Context) {
	gdb, _ := c.MustGet("databaseConn").(*gorm.DB)
	profile := Profile{Name: "Gopher"}
	gdb.Create(&profile)
}
