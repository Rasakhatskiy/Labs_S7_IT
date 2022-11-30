package main

import (
	"DBMS/database"
	"testing"
)

const (
	testDBName = "test"
)

func TestCreateDatabase(t *testing.T) {
	err := database.CreateDatabase(testDBName)

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestReadDatabasePaths(t *testing.T) {
	result := database.ReadDatabasesPaths()

	for _, json := range result {
		if json.Name == testDBName {
			return
		}
	}

	t.Fatal("DB not found")
}

func TestLoadDatabase(t *testing.T) {
	db, err := database.LoadDatabase(testDBName)

	if err != nil {
		t.Fatal(err.Error())
	}

	if db.Name != testDBName {
		t.Fatal("loaded wrong db")
	}
}

func TestAddTable(t *testing.T) {
	db, err := database.LoadDatabase(testDBName)
	if err != nil {
		t.Fatal(err.Error())
	}

	table := database.Table{
		Name: "dogs",
		Types: []string{
			database.TypeIntegerTS,
			database.TypeStringTS,
			database.TypeIntegerTS,
			database.TypeStringRangeTS,
			database.TypeHTMLTS,
		},
		Headers: []string{
			"id",
			"name",
			"owner_id",
			"range",
			"page",
		},
		Values: nil,
	}

	err = db.AddTable(table)

	if err != nil {
		t.Fatal(err.Error())
	}
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	if true {
		t.Fatalf("E")
	}
}
