package main

import (
	"DBMS/database"
	"DBMS/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/fs"
	"net/http"
	"strconv"
)

func main() {
	//createTestJson()

	e := echo.New()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		_ = c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": err.Error(),
		})
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// DATABASES
	e.GET("/databases", getDatabases)
	e.POST("/databases/new_database", createDB)

	// TABLES
	e.GET("/databases/:name", getTables)
	e.GET("/databases/:name/:table", getTable)
	e.POST("/databases/:name/new_table", addTable)
	e.GET("/databases/:name/join_tables", getJoinTablesData)
	e.GET("/databases/:name/joined_tables", getJoinedTables)

	// ROW
	e.POST("/databases/:name/:table/new_row", addRow)
	e.POST("/databases/:name/:table/:rowID", editRow)
	e.DELETE("/databases/:name/:table/:rowID", deleteRow)

	e.Logger.Fatal(e.Start(":1323"))
}

// e.POST("/databases/new_database", createDB)
func createDB(c echo.Context) error {
	data := new(string)
	err := c.Bind(data)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(*data)

	err = database.CreateDatabase(*data)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusCreated, data)
}

// e.POST("/databases/:name/new_table", addTable)
func addTable(c echo.Context) error {
	databaseName := c.Param("name")
	data := new(database.TableJSONValues)

	err := c.Bind(data)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(*data)

	db, err := database.LoadDatabase(databaseName)
	if err != nil {
		switch err.(type) {
		case *utils.DatabaseNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return err
		}
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
		return c.String(http.StatusBadRequest, "bad request")
	}

	err = db.SaveDatabase()
	if err != nil {
		return c.String(http.StatusInternalServerError, "can not save database")
	}

	return c.JSON(http.StatusCreated, data)
}

// e.POST("/databases/:name/:table/new_row", addRow)
func addRow(c echo.Context) error {
	data := new([]string)
	err := c.Bind(data)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(*data)

	databaseName := c.Param("name")
	tableName := c.Param("table")

	db, err := database.LoadDatabase(databaseName)
	if err != nil {
		switch err.(type) {
		case *utils.DatabaseNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	table, err := db.GetTable(tableName)

	if err != nil {
		switch err.(type) {
		case *utils.TableNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	err = table.AddRecord(*data)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request: "+err.Error())
	}

	err = db.SaveDatabase()
	if err != nil {
		return c.String(http.StatusInternalServerError, "can not save database")
	}

	return c.JSON(http.StatusCreated, data)
}

// e.DELETE("/databases/:name/:table/:rowID", deleteRow)
func deleteRow(c echo.Context) error {
	databaseName := c.Param("name")
	tableName := c.Param("table")
	index, _ := strconv.Atoi(c.Param("rowID"))

	db, err := database.LoadDatabase(databaseName)
	if err != nil {
		switch err.(type) {
		case *utils.DatabaseNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	table, err := db.GetTable(tableName)

	if err != nil {
		switch err.(type) {
		case *utils.TableNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	err = table.DeleteRecord(index)
	if err != nil {
		switch err.(type) {
		case *utils.TableNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	err = db.SaveDatabase()
	if err != nil {
		return c.String(http.StatusInternalServerError, "can not save database")
	}

	return c.String(http.StatusOK, "deleted")
}

// e.PUT("/databases/:name/:table/:rowID", editRow)
func editRow(c echo.Context) error {
	data := new([]string)
	err := c.Bind(data)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(*data)

	databaseName := c.Param("name")
	tableName := c.Param("table")
	index, _ := strconv.Atoi(c.Param("rowID"))

	db, err := database.LoadDatabase(databaseName)
	if err != nil {
		switch err.(type) {
		case *utils.DatabaseNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	table, err := db.GetTable(tableName)

	if err != nil {
		switch err.(type) {
		case *utils.TableNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	err = table.UpdateRecord(index, *data)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request: "+err.Error())
	}

	err = db.SaveDatabase()
	if err != nil {
		return c.String(http.StatusInternalServerError, "can not save database")
	}

	return c.String(http.StatusOK, "modified")
}

func getDatabases(c echo.Context) error {
	response := database.ReadDatabasesPaths()
	return c.JSON(http.StatusOK, response)
}

func getTables(c echo.Context) error {
	databaseName := c.Param("name")
	db, err := database.LoadDatabase(databaseName)
	if err != nil {
		switch err.(type) {
		case *fs.PathError:
			return c.JSON(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	tables := db.GetTablesList()

	return c.JSON(http.StatusOK, tables)
}

func tableToJson(table *database.Table) *database.TableJSONValues {
	var headers []database.TableHeaderJSON
	for i, _ := range table.Headers {
		headers = append(headers, database.TableHeaderJSON{
			Name: table.Headers[i],
			Type: table.Types[i],
		})
	}

	interValues := make([][]interface{}, len(table.Values))
	for i, _ := range interValues {
		interValues[i] = make([]interface{}, len(table.Values[i]))
		for j, _ := range interValues[i] {
			interValues[i][j] = table.Values[i][j].Value()
		}
	}

	return &database.TableJSONValues{
		Name:    table.Name,
		Headers: headers,
		Values:  interValues,
	}
}

func getTable(c echo.Context) error {
	databaseName := c.Param("name")
	tableName := c.Param("table")

	db, err := database.LoadDatabase(databaseName)
	if err != nil {
		switch err.(type) {
		case *utils.DatabaseNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	table, err := db.GetTable(tableName)
	if err != nil {
		switch err.(type) {
		case *utils.TableNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	return c.JSON(http.StatusOK, tableToJson(table))
}

func getJoinedTable(c echo.Context) error {
	databaseName := c.Param("name")
	tableName := c.Param("table")

	db, err := database.LoadDatabase(databaseName)
	if err != nil {
		switch err.(type) {
		case *utils.DatabaseNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	table, err := db.GetTable(tableName)
	if err != nil {
		switch err.(type) {
		case *utils.TableNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	return c.JSON(http.StatusOK, tableToJson(table))
}

func getJoinTablesData(c echo.Context) error {
	databaseName := c.Param("name")
	db, err := database.LoadDatabase(databaseName)
	if err != nil {
		switch err.(type) {
		case *fs.PathError:
			return c.JSON(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	jsonInfo := db.GetJSONInfo()
	return c.JSON(http.StatusOK, jsonInfo)
}

func getJoinedTables(c echo.Context) error {
	databaseName := c.Param("name")
	db, err := database.LoadDatabase(databaseName)
	if err != nil {
		switch err.(type) {
		case *fs.PathError:
			return c.JSON(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	t1 := c.QueryParam("t1")
	table1, err := db.GetTable(t1)
	if err != nil {
		switch err.(type) {
		case *utils.TableNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	c1 := c.QueryParam("c1")

	t2 := c.QueryParam("t2")
	table2, err := db.GetTable(t2)
	if err != nil {
		switch err.(type) {
		case *utils.TableNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return err
		}
	}

	c2 := c.QueryParam("c2")

	resTable, err := database.JoinTables(*table1, *table2, c1, c2)

	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, tableToJson(resTable))
}
