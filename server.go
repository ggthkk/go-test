package todoServer

import (
	"fmt"
	"log"
	"os"
	"todoServer/db"
	"todoServer/route"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
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

	Startbot(r, resource)
	route.RouteRanking(r, resource)
	r.Run(":27017")
}

func Startbot(r *gin.Engine, resource *db.Resource) {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			CreateMessage(resource, update.Message.Text)
		}
	}

}

func CreateMessage(res *db.Resource, message string) {
	ctx, _ := db.InitContext()
	fmt.Println("update botmessage...")
	fillter := bson.M{"name": "got"}

	setUpdate := bson.M{"$set": bson.M{"description": message}}

	_, err := res.DB.Collection("ranking").UpdateOne(ctx, fillter, setUpdate)

	if err != nil {
		fmt.Println("update botmessage error")
	}

}
