package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulreghupdm/tashub-project/internal/services"
)

func Register(c *gin.Context) {
	req := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	err := services.Register(req.Email, req.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
	req := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	access, refresh, err := services.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(200, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}
