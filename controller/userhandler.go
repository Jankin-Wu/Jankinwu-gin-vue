package controller

import (
	"Jankinwu-gin-vue/util"
	"github.com/gin-gonic/gin"
	"github.com/shuwenhe/utils"
	"net/http"
)

func Register(c *gin.Context) {
	// get the parameter
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	// Data verification
	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The phone num must be 11 digits!",
		})
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The password must be greater than 6 digits!",
		})
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	c.JSON(utils.NewSucc("Register success!", gin.H{
		"msg": "Register success!",
	}))
}
