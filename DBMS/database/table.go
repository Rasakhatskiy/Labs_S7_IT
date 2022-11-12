package database

import (
	"DBMS/utils"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Table struct {
	Name    string
	Types   []string
	Headers []string
	Values  [][]DBType
}

func CreateTable(name string, types []reflect.Type, headers []string) Table {
	var stringTypes []string
	for _, t := range types {
		stringTypes = append(stringTypes, t.String())
	}

	return Table{
		Name:    name,
		Headers: headers,
		Types:   stringTypes,
		Values:  nil,
	}
}

func (t *Table) hasColumn(name string) bool {
	return utils.Contains(t.Headers, name)
}

func JoinTables(t1, t2 Table, column1, column2 string) (*Table, error) {
	if !t1.hasColumn(column1) {
		return nil, &utils.ColumnDoesntExistsError{TableName: t1.Name, ColumnName: column1}
	}
	if !t2.hasColumn(column2) {
		return nil, &utils.ColumnDoesntExistsError{TableName: t2.Name, ColumnName: column2}
	}

	column1Index, err := utils.Find(t1.Headers, column1)
	column2Index, err := utils.Find(t2.Headers, column2)
	if err != nil {
		return nil, err
	}

	types2withoutColumn := utils.RemoveIndex(t2.Types, column2Index)
	headers2withoutColumn := utils.RemoveIndex(t2.Headers, column2Index)

	var values [][]DBType

	for _, row1 := range t1.Values {
		for _, row2 := range t2.Values {
			if row1[column1Index] == row2[column2Index] {
				values = append(values, row1, utils.RemoveIndex(row2, column2Index))
			}
		}
	}

	table := Table{
		Name:    fmt.Sprintf("%s_join_%s", t1.Name, t2.Name),
		Types:   append(t1.Types, types2withoutColumn...),
		Headers: append(t1.Headers, headers2withoutColumn...),
		Values:  values,
	}

	return &table, nil
}

func (t *Table) CreateRecord(values []DBType) error {
	for i, value := range values {
		if reflect.TypeOf(value).String() != t.Types[i] {
			return errors.New("Types mismatch")
		}
	}
	t.Values = append(t.Values, values)
	return nil
}

func (t *Table) ReadRecord(i int) []DBType {
	return t.Values[i]
}

func (t *Table) UpdateRecord(id int, dataStr []string) error {
	if id < 0 || id > len(t.Values) {
		return &utils.InvalidIndexError{Id: id}
	}

	values, err := t.parseDataTypes(dataStr)

	if err != nil {
		return err
	}

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

func (t *Table) parseDataTypes(data []string) ([]DBType, error) {
	var result []DBType
	typeOffset := 0
	for i := 0; i < len(data); i++ {
		str := data[i]
		var value DBType

		switch t.Types[i-typeOffset] {

		case TypeIntegerTS:
			parsed, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return nil, &utils.InvalidDataError{Data: str, TypeStr: TypeIntegerTS}
			}
			value = &TypeInteger{Val: parsed}

		case TypeRealTS:
			parsed, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return nil, &utils.InvalidDataError{Data: str, TypeStr: TypeRealTS}
			}
			value = &TypeReal{Val: parsed}

		case TypeStringTS:
			value = &TypeString{Val: str}

		case TypeStringRangeTS:
			strRange := TypeStringRange{}
			s1, s2 := str, data[i+1]

			err := strRange.Validate(s1, s2)
			if err != nil {
				return nil, &utils.InvalidDataError{Data: str, TypeStr: TypeStringRangeTS}
			}
			i++
			typeOffset++
			value = &strRange

		case TypeHTMLTS:
			html := TypeHTML{}
			err := html.Validate(str)
			if err != nil {
				return nil, &utils.InvalidDataError{Data: str, TypeStr: TypeHTMLTS}
			}
			value = &html

		case TypeCharTS:
			if len(str) != 1 {
				return nil, &utils.InvalidDataError{Data: str, TypeStr: TypeCharTS}
			}
			r := []rune(str)
			value = &TypeChar{Val: r[0]}
		}

		result = append(result, value)
	}

	return result, nil
}

func (t *Table) AddRecord(data []string) error {
	row, err := t.parseDataTypes(data)
	if err != nil {
		return err
	}
	t.Values = append(t.Values, row)
	return nil
}
