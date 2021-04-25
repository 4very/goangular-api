package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sommea/goangular-api/config"
	routes "github.com/sommea/goangular-api/routes"
)

func main() {
	config.Connect()

	r := gin.Default()

	routes.Routes(r)

	r.Run()
}
