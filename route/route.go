package route

import (
	"todoServer/controller"

	"github.com/gin-gonic/gin"
)

func RouteRanking(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/ranking/:name", controller.GetRanking())
	api.POST("/ranking", controller.SaveRanking())
}
