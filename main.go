package main

import (
	"github.com/gildemberg-santos/crudmondodb/mongodb"
)

func main() {
	// ...
	mongo := mongodb.Connection{}
	mongo.Connect()
}
