package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

const DBNamesListJson string = "databases.json"

type Database struct {
	name     string
	filePath string
	tables   []Table
}

func readDatabasesPaths() map[string]string {
	data, err := ioutil.ReadFile(DBNamesListJson)
	if err != nil {
		log.Fatal("Can not read databases path file")
	}

	var databasesPaths map[string]string
	err = json.Unmarshal(data, &databasesPaths)
	if err != nil {
		log.Fatal("Can not unmarshal databases path file json 😭")
	}
	return databasesPaths
}

func LoadDatabase(name string) (*Database, error) {
	//paths := readDatabasesPaths()
	//path, ok := paths[name]
	//log.Println(path)
	//if !ok {
	//	return nil, errors.New("database does not exist 😭")
	//}

	// todo read db from file
	return createTempDB(), nil
}

func createTempDB() *Database {
	typesList := []reflect.Type{reflect.TypeOf(TypeInteger{}), reflect.TypeOf(TypeInteger{}), reflect.TypeOf(TypeReal{}), reflect.TypeOf(TypeString{})}
	table := CreateTable("typical table name 😁", typesList)

	err := table.AddRecord([]DBType{
		TypeInteger{Val: 15},
		TypeInteger{Val: 15},
		TypeReal{Val: 47.50},
		TypeString{Val: "Brush"},
	})
	if err != nil {
		fmt.Printf(err.Error())
	}

	t2 := table
	t2.name = "WOW, new table!"

	t3 := table
	t3.name = "Another one bites the dust! 🙀"

	return &Database{
		name:     "amogus",
		filePath: "E",
		tables:   []Table{table, t2, t3},
	}
}

func (db *Database) GetTablesList() []string {
	var tables []string
	for _, table := range db.tables {
		tables = append(tables, table.name)
	}
	return tables
}

func CreateDatabase(name string) Database {
	defer log.Println("Created database: ", name)
	return Database{name: name}
}
