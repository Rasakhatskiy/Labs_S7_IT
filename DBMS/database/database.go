package database

import (
	"DBMS/utils"
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
	//path, ok := paths[Name]
	//log.Println(path)
	//if !ok {
	//	return nil, errors.New("database does not exist 😭")
	//}

	// todo read db from file
	return createTempDB(), nil
}

func createTempDB() *Database {
	typesList := []reflect.Type{reflect.TypeOf(TypeInteger{}), reflect.TypeOf(TypeInteger{}), reflect.TypeOf(TypeReal{}), reflect.TypeOf(TypeString{})}
	headers := []string{"header_1", "numero_2", "top_3_monki", "help_me"}
	table := CreateTable("table_01", typesList, headers)

	err := table.AddRecord([]DBType{
		TypeInteger{Val: 15},
		TypeInteger{Val: 15},
		TypeReal{Val: 47.50},
		TypeString{Val: "Brush"},
	})

	err = table.AddRecord([]DBType{
		TypeInteger{Val: 47},
		TypeInteger{Val: 74},
		TypeReal{Val: 11.22},
		TypeString{Val: "smth"},
	})

	if err != nil {
		fmt.Printf(err.Error())
	}

	t2 := table
	t2.Name = "table_02"

	t3 := table
	t3.Name = "table_03"

	return &Database{
		name:     "amogus",
		filePath: "E",
		tables:   []Table{table, t2, t3},
	}
}

func (db *Database) GetTablesList() []string {
	var tables []string
	for _, table := range db.tables {
		tables = append(tables, table.Name)
	}
	return tables
}

func (db *Database) GetTable(name string) (*Table, error) {
	for _, table := range db.tables {
		if table.Name == name {
			return &table, nil
		}
	}
	return nil, &utils.TableNotFoundError{TableName: name}
}

func CreateDatabase(name string) Database {
	defer log.Println("Created database: ", name)
	return Database{name: name}
}
