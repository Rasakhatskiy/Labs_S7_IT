package database

import "log"

type Database struct {
	name   string
	tables []Table
}

func CreateDatabase(name string) Database {
	defer log.Println("Created database: ", name)
	return Database{name: name}
}
