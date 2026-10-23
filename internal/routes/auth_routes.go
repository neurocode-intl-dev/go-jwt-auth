package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulreghupdm/tashub-project/internal/handlers"
)

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", handlers.Login)
	rg.POST("/register", handlers.Register)
}
