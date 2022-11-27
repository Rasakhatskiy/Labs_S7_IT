package database

import (
	"DBMS/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

const DBNamesListJson string = "databases.json"

type Database struct {
	name   string
	tables []Table
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
		name:   "amogus",
		tables: []Table{table, t2, t3},
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
	for i := range db.tables {
		if db.tables[i].Name == name {
			return &db.tables[i], nil
		}
	}
	return nil, &utils.TableNotFoundError{TableName: name}
}

func (db *Database) AddTable(table Table) error {
	db.tables = append(db.tables, table)
	return nil
}

func CreateDatabase(name string) error {
	db := Database{
		name:   name,
		tables: nil,
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
	for _, table := range db.tables {
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
