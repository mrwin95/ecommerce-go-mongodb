package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignIn(c *gin.Context) {

	var login Login
	if err := c.ShouldBindJSON((&login)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if login.Email != "win" || login.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login Success"})
}
