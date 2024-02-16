package todoServer

import (
	"todoServer/db"
	"todoServer/route"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	println("Starting server...")

	r := gin.Default()
	db.CreateResource()

	resource, err := db.CreateResource()
	if err != nil {
		println("Error creating resource")
		return
	}

	Startbot(resource)
	route.RouteRanking(r, resource)
	r.Run(":27017")
}

func Startbot(resource *db.Resource) {
	println("Starting bot...")
}
