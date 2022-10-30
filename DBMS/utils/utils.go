package utils

import (
	"fmt"
)

func RemoveIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func Contains[T comparable](array []T, elem T) bool {
	for _, el := range array {
		if elem == el {
			return true
		}
	}
	return false
}

func Find[T comparable](array []T, elem T) (int, error) {
	for i, el := range array {
		if elem == el {
			return i, nil
		}
	}
	return -1, &ItemNotFoundError{}
}

type TableNotFoundError struct {
	TableName string
}

func (m *TableNotFoundError) Error() string {
	return "table not found: " + m.TableName
}

type DatabaseNotFoundError struct {
	DatabaseName string
}

func (m *DatabaseNotFoundError) Error() string {
	return "database not found: " + m.DatabaseName
}

type InvalidIndexError struct {
	Id int
}

func (m *InvalidIndexError) Error() string {
	return fmt.Sprintf("invalid index: %d", m.Id)
}

type ColumnDoesntExistsError struct {
	TableName  string
	ColumnName string
}

func (m *ColumnDoesntExistsError) Error() string {
	return fmt.Sprintf("column '%s' does not exists in table '%s'", m.ColumnName, m.TableName)
}

type ItemNotFoundError struct {
}

func (m *ItemNotFoundError) Error() string {
	return fmt.Sprintf("item not found")
}

type InvalidDataError struct {
	Data    string
	TypeStr string
}

func (m *InvalidDataError) Error() string {
	return fmt.Sprintf("value '%s' is not of a type '%s'", m.Data, m.TypeStr)
}
