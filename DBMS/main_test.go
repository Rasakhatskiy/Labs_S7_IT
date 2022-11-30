package main

import (
	"DBMS/database"
	"testing"
)

func createTestDB() *database.Database {
	return &database.Database{
		Name: "Name",
		Tables: []database.Table{
			{
				Name:    "",
				Types:   nil,
				Headers: nil,
				Values:  nil,
			},
		},
	}
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	if true {
		t.Fatalf("E")
	}
}
