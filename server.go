package todoServer

import (
	"todoServer/route"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	println("Starting server...")

	r := gin.Default()

	route.RouteRanking(r)
	r.Run(":8892")
}
