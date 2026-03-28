package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Login(c *gin.Context) {
	var input loginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	if input.Username == "admin" && input.Password == "admin" {
		c.JSON(http.StatusOK, gin.H{
			"role":     "admin",
			"username": "admin",
		})
		return
	}

	if input.Username == "user" && input.Password == "user" {
		c.JSON(http.StatusOK, gin.H{
			"role":     "user",
			"username": "user",
		})
		return
	}

	newErrorResponse(c, http.StatusUnauthorized, "invalid credentials")
}
