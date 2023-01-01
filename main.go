package main

import "github.com/gildemberg-santos/crudmondodb/pkg"

func main() {
	conn := pkg.MongoDB{
		SettingLog: true,
	}

	conn.Connect()
	data := map[string]interface{}{
		"test": "test_01",
	}
	filter := map[string]interface{}{
		"test": "test",
	}
	// conn.InsertOne(data, "test", "test")
	// conn.InsertMany([]interface{}{data}, "test", "test")
	conn.UpdateOne(data, filter, "test", "test")
}
