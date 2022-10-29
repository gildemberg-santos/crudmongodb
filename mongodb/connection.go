package mongodb

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Connection struct {
	Conn string
}

func (c *Connection) Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	c.Conn = os.Getenv("MONGO_CONNECTION")
}
