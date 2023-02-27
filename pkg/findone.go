package pkg

func (m *MongoDB) FindOne(filter interface{}, collection string, document string) interface{} {
	client, ctx, _, errConnect := m.Connect()
	if errConnect != nil {
		m.Log("Error Connection MongoDB: " + errConnect.Error())
	}
	defer client.Disconnect(ctx)

	collectionMongo := client.Database(collection).Collection(document)
	result, err := collectionMongo.FindOne(ctx, filter).DecodeBytes()
	if err != nil {
		m.Log("Error FindOne MongoDB: " + err.Error())
	}

	m.Log("FindOne MongoDB: Success")
	return result
}
