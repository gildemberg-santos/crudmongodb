package pkg

func (m *MongoDB) InsertOne(insert interface{}, collection string, document string) {
	client, ctx, _, errConnect := m.Connect()
	if errConnect != nil {
		m.Log("Error Connection MongoDB: " + errConnect.Error())
	}
	defer client.Disconnect(ctx)

	collectionMongo := client.Database(collection).Collection(document)
	_, errInsert := collectionMongo.InsertOne(ctx, insert)
	if errInsert != nil {
		m.Log("Error Insert MongoDB: " + errInsert.Error())
	}

	m.Log("Insert MongoDB: Success")
}
