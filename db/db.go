package db

import (
	"context"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout           = 30
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

type Resource struct {
	DB *mongo.Database
}

func CreateResource() (*Resource, error) {
	_ = godotenv.Load()
	var err error
	var client *mongo.Client
	var dbName string
	var connectionURI string
	dbName = os.Getenv("MONGODB_DB_NAME")
	connectionURI = os.Getenv("DATABASE")
	client, err = mongo.NewClient(
		options.Client().ApplyURI(connectionURI),
		options.Client().SetMinPoolSize(1),
		options.Client().SetMaxPoolSize(2),
	)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	_ = client.Connect(ctx)
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	color.Green("Connect database successfully")
	color.Green(connectionURI)
	return &Resource{DB: client.Database(dbName)}, nil
}

func (r *Resource) Close() {
	ctx, cancel := InitContext()
	defer cancel()

	if err := r.DB.Client().Disconnect(ctx); err != nil {
		color.Red("Close connection falure, Something wrong...")
		return
	}

	color.Cyan("Close connection successfully")
}

func InitContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	return ctx, cancel
}
