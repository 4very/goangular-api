package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/sommea/goangular-api/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/players", controllers.GetAllPlayers)
	router.GET("/getplayer", controllers.GetPlayer)
	router.POST("/players", controllers.CreatePlayer)
	router.GET("/guilds", controllers.GetAllGuilds)
	router.POST("/guilds", controllers.CreateGuild)
	router.POST("/report", controllers.CreateReport)
	router.GET("/report", controllers.GetAllReports)
	router.GET("/fight", controllers.GetAllFights)

}
func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
}
