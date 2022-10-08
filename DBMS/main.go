package main

import (
	"DBMS/database"
	"fmt"
	"reflect"
)

var gCurrentDB database.Database

func main() {

	suck := []reflect.Type{reflect.TypeOf(database.TypeInteger{}), reflect.TypeOf(database.TypeInteger{}), reflect.TypeOf(database.TypeReal{}), reflect.TypeOf(database.TypeString{})}
	table := database.CreateTable("a", suck)

	err := table.AddRecord([]database.DBType{
		database.TypeInteger{Val: 15},
		database.TypeInteger{Val: 15},
		database.TypeReal{Val: 47.50},
		database.TypeString{Val: "Brush"},
	})
	if err != nil {
		fmt.Printf(err.Error())
	}

}
