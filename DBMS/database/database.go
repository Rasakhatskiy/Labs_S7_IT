package database

import "log"

var gDatabases []*Database

type Database struct {
	name     string
	filePath string
	tables   []Table
}

func loadDatabases() {

}

func CreateDatabase(name string) Database {
	defer log.Println("Created database: ", name)
	return Database{name: name}
}
