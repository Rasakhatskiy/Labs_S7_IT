package database

import (
	"DBMS/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

const DBNamesListJson string = "databases.json"

type Database struct {
	Name   string
	Tables []Table
}

type DBPathJSON struct {
	Name string `json:"name"`
}

func ReadDatabasesPaths() []DBPathJSON {
	jsonFile, err := os.Open(DBNamesListJson)
	if err != nil {
		log.Fatal("Can not read databases path file")
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Can not read json file")
	}

	var databases []DBPathJSON

	if len(byteValue) == 0 {
		return databases
	}

	err = json.Unmarshal(byteValue, &databases)
	if err != nil {
		log.Fatal("Can not unmarshal databases path file json 😭")
	}

	return databases
}

func SaveDatabasesPaths(data []DBPathJSON) {
	file, err := json.MarshalIndent(data, "", "")

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(DBNamesListJson, file, 0644)
}

//func LoadDatabase(name string) (*Database, error) {
//	//paths := readDatabasesPaths()
//	//path, ok := paths[Name]
//	//log.Println(path)
//	//if !ok {
//	//	return nil, errors.New("database does not exist 😭")
//	//}
//
//	// todo read db from file
//	return createTempDB(), nil
//}

func createTempDB() *Database {
	typesList := []reflect.Type{reflect.TypeOf(TypeInteger{}), reflect.TypeOf(TypeInteger{}), reflect.TypeOf(TypeReal{}), reflect.TypeOf(TypeString{})}
	headers := []string{"header_1", "numero_2", "top_3_monki", "help_me"}
	table := CreateTable("table_01", typesList, headers)

	err := table.CreateRecord([]DBType{
		TypeInteger{Val: 15},
		TypeInteger{Val: 15},
		TypeReal{Val: 47.50},
		TypeString{Val: "Brush"},
	})

	err = table.CreateRecord([]DBType{
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
		Name:   "amogus",
		Tables: []Table{table, t2, t3},
	}
}

func (db *Database) GetTablesList() []string {
	var tables []string
	for _, table := range db.Tables {
		tables = append(tables, table.Name)
	}
	return tables
}

func (db *Database) GetTable(name string) (*Table, error) {
	for i := range db.Tables {
		if db.Tables[i].Name == name {
			return &db.Tables[i], nil
		}
	}
	return nil, &utils.TableNotFoundError{TableName: name}
}

func (db *Database) DeleteTable(name string) error {
	for i := range db.Tables {
		if db.Tables[i].Name == name {
			utils.RemoveIndex(db.Tables, i)
			return nil
		}
	}
	return &utils.TableNotFoundError{TableName: name}
}

func (db *Database) AddTable(table Table) error {
	db.Tables = append(db.Tables, table)
	return nil
}

func CreateDatabase(name string) error {
	db := Database{
		Name:   name,
		Tables: nil,
	}

	pathEntry := DBPathJSON{
		Name: name,
	}

	dbpaths := ReadDatabasesPaths()
	if !utils.Contains(dbpaths, pathEntry) {
		dbpaths = append(dbpaths, pathEntry)
	}
	SaveDatabasesPaths(dbpaths)

	err := db.SaveDatabase()
	return err
}

func (db *Database) GetJSONInfo() DatabaseInfoJSON {
	var infoJSON DatabaseInfoJSON
	for _, table := range db.Tables {
		var tableSJON TableJSON
		tableSJON.Name = table.Name
		for i := range table.Headers {
			tableSJON.Headers = append(tableSJON.Headers, TableHeaderJSON{
				Name: table.Headers[i],
				Type: table.Types[i],
			})
		}
		infoJSON.Tables = append(infoJSON.Tables, tableSJON)
	}
	return infoJSON
}

func DeleteDatabase(name string) error {
	databasesPaths := ReadDatabasesPaths()
	pathEntry := DBPathJSON{
		Name: name,
	}
	if !utils.Contains(databasesPaths, pathEntry) {
		return errors.New("no such database")
	}

	i, _ := utils.Find(databasesPaths, pathEntry)
	databasesPaths = utils.RemoveIndex(databasesPaths, i)
	SaveDatabasesPaths(databasesPaths)
	return nil
}
