package main

import (
	"./database"
)

func main() {
	database.Connect("mongodb://admin:i3540core@ds143242.mlab.com:43242/sweet-cafe", "sweet-cafe")
}
