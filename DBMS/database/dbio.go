package database

import (
	"encoding/gob"
	"fmt"
	"os"
)

func registerGOB() {
	gob.Register(TypeInteger{Val: 0})
	gob.Register(TypeReal{Val: 0})
	gob.Register(TypeChar{Val: 'A'})
	gob.Register(TypeString{Val: "A"})
	gob.Register(TypeHTML{Val: "E"})
	gob.Register(TypeStringRange{Val: []string{"A", "Z"}})
}

func (db *Database) SaveDatabase() error {
	registerGOB()

	f, err := os.Create(fmt.Sprintf("%s.gob", db.filePath))
	if err != nil {
		return err
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	if err := enc.Encode(db.name); err != nil {
		return err
	}

	if err := enc.Encode(db.tables); err != nil {
		return err
	}

	return nil
}

func LoadDatabase(filePath string) (*Database, error) {
	registerGOB()

	f, err := os.Open(fmt.Sprintf("%s.gob", filePath))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var db Database

	dec := gob.NewDecoder(f)
	if err := dec.Decode(&db.name); err != nil {
		return nil, err
	}

	if err := dec.Decode(&db.tables); err != nil {
		return nil, err
	}

	return &db, nil
}