package controller

import "github.com/gin-gonic/gin"

func GetRanking() func(c *gin.Context) {
	return func(c *gin.Context) {

		name := c.Params.ByName("name")
		c.JSON(200, gin.H{
			"name": name,
		})
	}
}

func SaveRanking() func(c *gin.Context) {

	type infobody struct {
		Name     string `json:"name"`
		Lastname string `json:"lastname" binding:"required"`
	}
	return func(c *gin.Context) {
		var data infobody

		if err := c.Bind(&data); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, data)
	}
}
