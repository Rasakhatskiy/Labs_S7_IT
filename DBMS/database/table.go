package database

import (
	"errors"
	"fmt"
	"reflect"
)

type Table struct {
	types   []reflect.Type
	records []Record
}

func CreateTable(types []reflect.Type) Table {
	return Table{
		types:   types,
		records: nil,
	}
}

func (t *Table) AddRecord(values []DBType) error {
	for i, value := range values {
		if reflect.TypeOf(value) != t.types[i] {
			return errors.New("types mismatch")
		}
	}
	t.records = append(t.records, Record{values: values})

	for _, val := range t.records[len(t.records)-1].values {
		fmt.Println(val.Value())
	}

	return nil
}
