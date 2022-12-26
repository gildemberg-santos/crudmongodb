package pkg

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Conn       string
	Panic      bool
	SettingLog bool
}

func (m *MongoDB) Connect() (client *mongo.Client, ctx context.Context, fctx context.CancelFunc, err error) {
	godotenv.Load()
	m.Conn = os.Getenv("MONGO_CONNECTION")

	clientOptions := options.Client().ApplyURI(m.Conn)
	client, err = mongo.NewClient(clientOptions)
	if err != nil {
		m.PanicLog(err)
	}

	ctx, fctx = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		m.PanicLog(err)
	}
	return
}
