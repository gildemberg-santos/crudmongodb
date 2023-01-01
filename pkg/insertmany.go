package pkg

func (m *MongoDB) InsertMany(inserts []interface{}, collection string, document string) {
	client, ctx, _, errConnect := m.Connect()
	if errConnect != nil {
		m.Log("Error Connection MongoDB: " + errConnect.Error())
	}
	defer client.Disconnect(ctx)

	collectionMongo := client.Database(collection).Collection(document)
	_, errInsert := collectionMongo.InsertMany(ctx, inserts)
	if errInsert != nil {
		m.Log("Error Insert MongoDB: " + errInsert.Error())
	}

	m.Log("Insert MongoDB: Success")
}
