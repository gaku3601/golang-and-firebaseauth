package controllers

import (
	"github.com/gaku3601/golang-and-firebaseauth/models"
	auth "github.com/gaku3601/golang-and-firebaseauth/util"
	"github.com/gin-gonic/gin"
)

//GetProfiles get method
func GetProfiles(r *gin.RouterGroup) {
	r.GET("/", auth.Auth(func(c *gin.Context) { // TODO: AUTH使えるか検証する
		c.String(200, "Hello,World!")
	}))
}

//PostProfiles get method
func PostProfiles(r *gin.RouterGroup) {
	r.POST("/", func(c *gin.Context) {
		models.CreateProfile(c)
		c.String(200, "Hello,World!")
	})
}
