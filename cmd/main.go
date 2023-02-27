package main

import (
	"fmt"

	"github.com/gildemberg-santos/crudmongodb/pkg"
)

func main() {
	conn := pkg.MongoDB{
		SettingLog: true,
	}

	conn.Connect()
	// data := map[string]interface{}{
	// 	"test": "test",
	// }
	filter := map[string]interface{}{
		"test": "test",
	}
	// conn.InsertOne(data, "test", "test")
	// conn.InsertMany([]interface{}{data}, "test", "test")
	// conn.UpdateOne(data, filter, "test", "test")
	fmt.Println(conn.FindOne(filter, "test", "test"))
}
