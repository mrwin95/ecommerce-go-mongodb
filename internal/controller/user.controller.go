package controller

import (
	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(200, gin.H{
		"message": `Get User By` + id,
	})
}

func CreateUser(c *gin.Context) {

	// var user c.PostForm("user")
	// if user nil {
	// 	c.JSON(400, gin.H{
	// 		"message": "User is required",
	// 	})
	// 	return
	// }
	c.JSON(200, gin.H{
		"message": "Create User",
	})
}
