package serverImpl

import (
	"DBMS/database"
	"DBMS/genserver"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

const (
	StatusBadRequest = "bad request"
	StatusDeleted    = "deleted"
	StatusCantSave   = "cant save database"
	StatusModifies   = "modified"
)

// SetupHandler creates echo server, registers handlers and starts server
func SetupHandler() {
	var myApi MyServerImpl
	e := echo.New()
	genserver.RegisterHandlers(e, &myApi)
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		_ = c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": err.Error(),
		})
	}
	e.Logger.Fatal(e.Start(":1323"))
}

func tableToJson(table *database.Table) *database.TableJSONValues {
	var headers []database.TableHeaderJSON
	for i := range table.Headers {
		headers = append(headers, database.TableHeaderJSON{
			Name: table.Headers[i],
			Type: table.Types[i],
		})
	}

	interValues := make([][]interface{}, len(table.Values))
	for i := range interValues {
		interValues[i] = make([]interface{}, len(table.Values[i]))
		for j := range interValues[i] {
			interValues[i][j] = table.Values[i][j].Value()
		}
	}

	return &database.TableJSONValues{
		Name:    table.Name,
		Headers: headers,
		Values:  interValues,
	}
}

type MyServerImpl struct {
}

func (s *MyServerImpl) GetDatabases(ctx echo.Context) error {
	response := database.ReadDatabasesPaths()
	return ctx.JSON(http.StatusOK, response)
}

func (s *MyServerImpl) PostDatabases(ctx echo.Context) error {
	data := new(string)
	err := ctx.Bind(data)
	if err != nil {
		return ctx.String(http.StatusBadRequest, StatusBadRequest)
	}
	fmt.Println(*data)

	err = database.CreateDatabase(*data)
	if err != nil {
		return ctx.String(http.StatusBadRequest, StatusBadRequest)
	}

	return ctx.JSON(http.StatusCreated, data)
}

func (s *MyServerImpl) DeleteDatabasesDbName(ctx echo.Context, dbName string) error {
	err := database.DeleteDatabase(dbName)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, StatusDeleted)
}

func (s *MyServerImpl) PostDatabasesDbName(ctx echo.Context, dbName string) error {
	data := new(database.TableJSONValues)

	err := ctx.Bind(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, StatusBadRequest)
	}
	fmt.Println(*data)

	db, err := database.LoadDatabase(dbName)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	var headers []string
	var types []string

	for _, header := range data.Headers {
		headers = append(headers, header.Name)
		types = append(types, header.Type)
	}

	table := database.Table{
		Name:    data.Name,
		Types:   types,
		Headers: headers,
		Values:  nil,
	}

	err = db.AddTable(table)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, StatusBadRequest)
	}

	err = db.SaveDatabase()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, StatusCantSave)
	}

	return ctx.JSON(http.StatusCreated, data)
}

func (s *MyServerImpl) GetDatabasesDbNameJoinTables(ctx echo.Context, dbName string) error {
	db, err := database.LoadDatabase(dbName)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	jsonInfo := db.GetJSONInfo()
	return ctx.JSON(http.StatusOK, jsonInfo)
}

func (s *MyServerImpl) GetDatabasesDbNameJoinedTables(ctx echo.Context, dbName string, params genserver.GetDatabasesDbNameJoinedTablesParams) error {
	db, err := database.LoadDatabase(dbName)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	table1, err := db.GetTable(params.T1)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	table2, err := db.GetTable(params.T2)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	resTable, err := database.JoinTables(*table1, *table2, params.C1, params.C2)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, tableToJson(resTable))
}

func (s *MyServerImpl) DeleteDatabasesDbNameTableName(ctx echo.Context, dbName string, tableName string) error {
	db, err := database.LoadDatabase(dbName)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	err = db.DeleteTable(tableName)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	err = db.SaveDatabase()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, StatusCantSave)
	}

	return ctx.JSON(http.StatusOK, "deleted")
}

func (s *MyServerImpl) GetDatabasesDbNameTableName(ctx echo.Context, dbName string, tableName string) error {
	db, err := database.LoadDatabase(dbName)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	table, err := db.GetTable(tableName)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, tableToJson(table))
}

func (s *MyServerImpl) PostDatabasesDbNameTableName(ctx echo.Context, dbName string, tableName string) error {
	data := new([]string)
	err := ctx.Bind(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, StatusBadRequest)
	}

	db, err := database.LoadDatabase(dbName)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	table, err := db.GetTable(tableName)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	err = table.AddRecord(*data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, StatusBadRequest)
	}

	err = db.SaveDatabase()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, StatusCantSave)
	}

	return ctx.JSON(http.StatusCreated, data)
}

func (s *MyServerImpl) DeleteDatabasesDbNameTableNameRowId(ctx echo.Context, dbName string, tableName string, rowId int) error {
	db, err := database.LoadDatabase(dbName)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	table, err := db.GetTable(tableName)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	err = table.DeleteRecord(rowId)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	err = db.SaveDatabase()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, StatusCantSave)
	}

	return ctx.JSON(http.StatusOK, StatusDeleted)
}

func (s *MyServerImpl) PutDatabasesDbNameTableNameRowId(ctx echo.Context, dbName string, tableName string, rowId int) error {
	data := new([]string)
	err := ctx.Bind(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, StatusBadRequest)
	}

	db, err := database.LoadDatabase(dbName)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	table, err := db.GetTable(tableName)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	err = table.UpdateRecord(rowId, *data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, StatusBadRequest)
	}

	err = db.SaveDatabase()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, StatusCantSave)
	}

	return ctx.JSON(http.StatusOK, StatusModifies)
}
