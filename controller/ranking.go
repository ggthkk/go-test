package controller

import (
	"todoServer/db"
	"todoServer/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRanking(res *db.Resource) func(c *gin.Context) {
	var ranking model.Ranking
	return func(c *gin.Context) {

		code := c.Query("code")
		res.DB.Collection("ranking").FindOne(c, bson.M{"code": code}).Decode(&ranking)
		c.JSON(200, ranking)
	}
}

func GetRankingList(res *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var ranking []model.Ranking

		code := c.Query("code")
		curcer, _ := res.DB.Collection("ranking").Find(c, bson.M{"code": code})

		curcer.All(c, &ranking)

		c.JSON(200, ranking)
	}
}

func CreateRanking(res *db.Resource) func(c *gin.Context) {
	var ranking model.Ranking
	return func(c *gin.Context) {

		if err := c.Bind(&ranking); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		_, err := res.DB.Collection("ranking").InsertOne(c, ranking)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, ranking)
	}
}

func UpdateRanking(res *db.Resource) func(c *gin.Context) {
	var ranking model.Ranking
	return func(c *gin.Context) {
		if err := c.Bind(&ranking); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		name := c.Param("name")

		fillter := bson.M{"name": name}

		setUpdate := bson.M{"$set": ranking}

		_, err := res.DB.Collection("ranking").UpdateOne(c, fillter, setUpdate)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, ranking)
	}
}

func DeleteRanking(res *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("name")

		_, err := res.DB.Collection("ranking").DeleteOne(c, bson.M{"name": name})
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "success"})
	}
}
