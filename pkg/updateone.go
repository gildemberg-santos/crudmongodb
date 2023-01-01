package pkg

import "go.mongodb.org/mongo-driver/bson"

func (m *MongoDB) UpdateOne(upload interface{}, filter interface{}, collection string, document string) {
	client, ctx, _, errConnect := m.Connect()
	if errConnect != nil {
		m.Log("Error Connection MongoDB: " + errConnect.Error())
	}
	defer client.Disconnect(ctx)

	collectionMongo := client.Database(collection).Collection(document)
	_, errUpdate := collectionMongo.UpdateOne(ctx, filter, bson.M{"$set": upload})
	if errUpdate != nil {
		m.Log("Error Update MongoDB: " + errUpdate.Error())
	}

	m.Log("Update MongoDB: Success")
}
