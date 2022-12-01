package main

import (
	"DBMS/database"
	"testing"
)

const (
	testDBName = "test"
)

var testTable1 = database.Table{
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

var testTable2 = database.Table{
	Name: "owners",
	Types: []string{
		database.TypeIntegerTS,
		database.TypeStringTS,
	},
	Headers: []string{
		"id",
		"name",
	},
	Values: nil,
}

var testValues1 = [][]database.DBType{
	{
		database.TypeInteger{Val: 1},
		database.TypeString{Val: "bobrik"},
		database.TypeInteger{Val: 2},
		database.TypeStringRange{Val: []string{"aaaa", "bbbb"}},
		database.TypeHTML{Val: "<div> a </div>"},
	},
}

var testValues2 = [][]database.DBType{
	{
		database.TypeInteger{Val: 2},
		database.TypeString{Val: "Oleh"},
	},
}

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

	err = db.AddTable(testTable1)

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetTable(t *testing.T) {
	db, err := database.LoadDatabase(testDBName)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = db.AddTable(testTable1)
	if err != nil {
		t.Fatal(err.Error())
	}

	table, err := db.GetTable("dogs")
	if err != nil {
		t.Fatal(err.Error())
	}

	if table.Name != "dogs" || table.Headers[2] != "owner_id" {
		t.Fatalf("read wrong table")
	}
}

func TestJoinTables(t *testing.T) {
	testTable1.Values = testValues1
	testTable2.Values = testValues2

	joined, err := database.JoinTables(testTable1, testTable2, "owner_id", "id")
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(joined.Values) != 1 || len(joined.Values[0]) != 6 {
		t.Fatalf("wrong joined size")
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		if joined.Values[0][i].Value() != testValues1[0][i].Value() {
			t.Fatalf("wrong joined result %d", i)
		}
	}

	if joined.Values[0][5].Value() != testValues2[0][1].Value() {
		t.Fatalf("wrong joined result 5")
	}
}

func TestValidateStringRange_1(t *testing.T) {
	data := database.TypeStringRange{}
	err := data.Validate("arbol", "bruh momento")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestValidateStringRange_2(t *testing.T) {
	data := database.TypeStringRange{}
	err := data.Validate("bober", "bebra")
	if err == nil {
		t.Fatal("invalid validation")
	}
}

func TestValidateHTML_1(t *testing.T) {
	data := database.TypeHTML{}
	err := data.Validate("<div> <bruh> <momento> numero 2 </momento> </bruh> </div>")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestValidateHTML_2(t *testing.T) {
	data := database.TypeHTML{}
	err := data.Validate("<div> <bruh> <momento> numero 2 </momento> </div>")
	if err == nil {
		t.Fatal("invalid validation")
	}
}
