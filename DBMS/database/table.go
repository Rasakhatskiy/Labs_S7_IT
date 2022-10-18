package database

import (
	"DBMS/utils"
	"errors"
	"reflect"
)

type Table struct {
	Name    string
	Types   []reflect.Type
	Headers []string
	Values  [][]DBType
}

func CreateTable(name string, types []reflect.Type, headers []string) Table {
	return Table{
		Name:    name,
		Headers: headers,
		Types:   types,
		Values:  nil,
	}
}

func (t *Table) CreateRecord(values []DBType) error {
	for i, value := range values {
		if reflect.TypeOf(value) != t.Types[i] {
			return errors.New("Types mismatch")
		}
	}
	t.Values = append(t.Values, values)
	return nil
}

func (t *Table) ReadRecord(i int) []DBType {
	return t.Values[i]
}

func (t *Table) UpdateRecord(id int, values []DBType) error {
	if id < 0 || id > len(t.Values) {
		return &utils.InvalidIndexError{Id: id}
	}
	t.Values[id] = values
	return nil
}

func (t *Table) DeleteRecord(id int) error {
	if id < 0 || id > len(t.Values) {
		return &utils.InvalidIndexError{Id: id}
	}
	t.Values = utils.RemoveIndex(t.Values, id)
	return nil
}
