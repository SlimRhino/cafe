package database

import (
	"log"

	"github.com/globalsign/mgo"
)

//DB Ссылка для подключения к базе данных MongoDB
var DB *mgo.Database

// Connect to database
func Connect(server string, dbName string) {
	session, err := mgo.Dial(server)
	if err != nil {
		log.Fatal(err)
	}
	DB = session.DB(dbName)
	log.Println("connect to database", DB.Name, "succssesed")
}
