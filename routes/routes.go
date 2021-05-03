package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/sommea/goangular-api/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)

	router.POST("/player", controllers.CreatePlayer)
	router.GET("/player", controllers.GetPlayer)
	router.GET("/player/:path1", controllers.GetPlayerURL)

	router.POST("/guild", controllers.CreateGuild)
	router.GET("/guild/all", controllers.GetAllGuilds)
	router.GET("/guild", controllers.GetGuild)

	router.POST("/report", controllers.CreateReport)
	router.GET("/report/all", controllers.GetAllReports)

	router.GET("/fight/all", controllers.GetAllFights)

	router.GET("/com/data/all", controllers.GetAllComData)
	router.POST("/com", controllers.UpdateComData)

}
func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
}
