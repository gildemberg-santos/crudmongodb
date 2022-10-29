package mongodb

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Conn string
}

func (c *Connection) Connect() (client *mongo.Client, ctx context.Context, fctx context.CancelFunc, err error) {
	godotenv.Load()
	c.Conn = os.Getenv("MONGO_CONNECTION")

	clientOptions := options.Client().ApplyURI(c.Conn)
	client, _ = mongo.NewClient(clientOptions)

	ctx, fctx = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	return
}
