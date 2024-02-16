package route

import (
	"todoServer/controller"
	"todoServer/db"

	"github.com/gin-gonic/gin"
)

func RouteRanking(r *gin.Engine, resource *db.Resource) {
	api := r.Group("/api")

	api.GET("/ranking/:name", controller.GetRanking(resource))
	api.GET("/rankinglists", controller.GetRankingList(resource))
	api.POST("/createranking", controller.CreateRanking(resource))
	api.PATCH("/UpdateRanking/:name", controller.UpdateRanking(resource))
	api.DELETE("/DeleteRanking/:name", controller.DeleteRanking(resource))
}
