package routes

import (
	"github.com/atharvamahamuni/golang/projects/restaurant-management/controllers"
	"github.com/atharvamahamuni/golang/projects/restaurant-management/middleware"
	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", middleware.Authentication(), controllers.GetTables())
	incomingRoutes.GET("/tables/:table_id", controllers.GetTable())
	incomingRoutes.POST("/tables", controllers.CreateTable())
	incomingRoutes.PUT("/tables/:table_id", controllers.UpdateTable())
}
