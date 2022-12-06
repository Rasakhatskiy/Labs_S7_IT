package clientImpl

import (
	"DBMS/genclient"
	"context"
	"io"
	"net/http"
)

type MyClientImpl struct {
}

func (m *MyClientImpl) GetDatabases(ctx context.Context, reqEditors ...genclient.RequestEditorFn) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MyClientImpl) PostDatabasesWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...genclient.RequestEditorFn) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MyClientImpl) PostDatabases(ctx context.Context, body genclient.PostDatabasesJSONRequestBody, reqEditors ...genclient.RequestEditorFn) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MyClientImpl) DeleteDatabasesDbName(ctx context.Context, dbName string, reqEditors ...genclient.RequestEditorFn) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MyClientImpl) PostDatabasesDbNameWithBody(ctx context.Context, dbName string, contentType string, body io.Reader, reqEditors ...genclient.RequestEditorFn) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MyClientImpl) PostDatabasesDbName(ctx context.Context, dbName string, body genclient.PostDatabasesDbNameJSONRequestBody, reqEditors ...genclient.RequestEditorFn) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MyClientImpl) GetDatabasesDbNameJoinTables(ctx context.Context, dbName string, reqEditors ...genclient.RequestEditorFn) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MyClientImpl) GetDatabasesDbNameJoinedTables(ctx context.Context, dbName string, params *genclient.GetDatabasesDbNameJoinedTablesParams, reqEditors ...genclient.RequestEditorFn) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MyClientImpl) DeleteDatabasesDbNameTableName(ctx context.Context, dbName string, tableName string, reqEditors ...genclient.RequestEditorFn) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MyClientImpl) GetDatabasesDbNameTableName(ctx context.Context, dbName string, tableName string, reqEditors ...genclient.RequestEditorFn) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MyClientImpl) PostDatabasesDbNameTableNameWithBody(ctx context.Context, dbName string, tableName string, contentType string, body io.Reader, reqEditors ...genclient.RequestEditorFn) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MyClientImpl) PostDatabasesDbNameTableName(ctx context.Context, dbName string, tableName string, body genclient.PostDatabasesDbNameTableNameJSONRequestBody, reqEditors ...genclient.RequestEditorFn) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MyClientImpl) DeleteDatabasesDbNameTableNameRowId(ctx context.Context, dbName string, tableName string, rowId int, reqEditors ...genclient.RequestEditorFn) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MyClientImpl) PutDatabasesDbNameTableNameRowId(ctx context.Context, dbName string, tableName string, rowId int, reqEditors ...genclient.RequestEditorFn) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}
