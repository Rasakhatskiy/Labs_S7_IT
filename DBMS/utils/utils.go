package utils

import "fmt"

func RemoveIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
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
	return fmt.Sprintf("database not found: %d", m.Id)
}
