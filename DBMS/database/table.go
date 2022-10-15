package database

import (
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

func (t *Table) AddRecord(values []DBType) error {
	for i, value := range values {
		if reflect.TypeOf(value) != t.Types[i] {
			return errors.New("Types mismatch")
		}
	}
	t.Values = append(t.Values, values)

	//for _, val := range t.Values[len(t.Values)-1].values {
	//	fmt.Println(val.Value())
	//}

	return nil
}

//func (t *Table) GetRecord(i int) Record {
//	return t.Values[i]
//}
