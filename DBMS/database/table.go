package database

import (
	"errors"
	"reflect"
)

type Table struct {
	types   []reflect.Type
	records []Record
}

func createTable(types []reflect.Type) Table {
	return Table{
		types:   types,
		records: nil,
	}
}

func (t *Table) addRecord(values []DBType) (Record, error) {
	for i, value := range values {
		if reflect.TypeOf(value) != t.types[i] {
			return Record{}, errors.New("types mismatch")
		}
	}
	return Record{values: values}, nil
}
