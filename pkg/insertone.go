package pkg

import (
	"log"
)

func (c *Connection) InsertOne(insert interface{}, collection string, document string) {
	client, ctx, _, errConnect := c.Connect()
	if errConnect != nil {
		log.Println("Error Connection MongoDB: ", errConnect)
	}
	defer client.Disconnect(ctx)

	collectionMongo := client.Database(collection).Collection(document)
	_, errInsert := collectionMongo.InsertOne(ctx, insert)
	if errInsert != nil {
		log.Println("Error Insert MongoDB: ", errInsert)
	}
}
