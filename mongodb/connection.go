package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Conn string
}

func (c *Connection) Connect(client *mongo.Client, ctx context.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	c.Conn = os.Getenv("MONGO_CONNECTION")

	clientOptions := options.Client().ApplyURI(c.Conn)
	client, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Println("Error MongoDB: ", err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println("Error MongoDB: ", err)
	}
}
