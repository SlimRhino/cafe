package database

import (
	"log"
	"strings"

	"github.com/globalsign/mgo"
)

//DB Ссылка для подключения к базе данных MongoDB
var DB *mgo.Database

var productCollections []string

// Connect to database
func Connect(server string, dbName string) {
	session, err := mgo.Dial(server)
	if err != nil {
		log.Fatalln(err)
	}
	DB = session.DB(dbName)
	log.Println("connect to database", DB.Name, "succssesed")
}

//GetProductCollections сохраняет имена существующих коллекций продуктов для дальнейшего сравнения
func GetProductCollections() {
	tempCollections, err := DB.CollectionNames()
	if err != nil {
		log.Fatalln(err)
	}

	for i := range tempCollections {
		if !strings.Contains(tempCollections[i], "prod_") {
			tempCollections = append(tempCollections[:i], tempCollections[i+1:]...)
		}
	}
	productCollections = tempCollections
	log.Println("product collections are filled.", "Total:", len(productCollections))
}

//CollectionExist Проверяет коллекцию на наличие в списке существующих в базе
func CollectionExist(col string) bool {
	for _, pCol := range productCollections {
		if pCol == col {
			return true
		}
	}
	return false
}
