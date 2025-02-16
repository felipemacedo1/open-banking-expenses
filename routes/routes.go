package routes

import (
	"github.com/gin-gonic/gin"
	"open-banking-expenses/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/transactions", controllers.GetTransactions)

	return r
}
