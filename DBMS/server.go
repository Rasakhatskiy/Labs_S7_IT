package main

import (
	"DBMS/database"
	"DBMS/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/fs"
	"net/http"
	"strconv"
)

const BADREQUEST = "bad request"

// @Summary		Get a table
// @Description	get table by name
// @Tags			table
// @Accept			json
// @Produce		json
// @Param			name	path		string	false	"Database name"
// @Param			table	path		string	false	"Table name"
// @Success		200		{object}	database.TableJSONValues
// @Router			/databases/{name}/{table} [get]
func getTable(c echo.Context) error {
	databaseName := c.Param("name")
	tableName := c.Param("table")

	db, err := database.LoadDatabase(databaseName)
	if err != nil {
		switch err.(type) {
		case *utils.DatabaseNotFoundError:
			return c.JSON(http.StatusNotFound, err.Error())
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

// e.POST("/databases/new_database", createDB)
//
//	@Summary		Create a database
//	@Description	Create database with given name
//	@Tags			database
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	false	"Database name"
//	@Success		200		{object}	string
//	@Router			/databases/new_database [post]
func createDB(c echo.Context) error {
	data := new(string)
	err := c.Bind(data)
	if err != nil {
		return c.String(http.StatusBadRequest, BADREQUEST)
	}
	fmt.Println(*data)

	err = database.CreateDatabase(*data)
	if err != nil {
		return c.String(http.StatusBadRequest, BADREQUEST)
	}

	return c.JSON(http.StatusCreated, data)
}

// e.POST("/databases/:name/new_table", addTable)
//
//	@Summary		Create new table
//	@Description	Create new table by name
//	@Tags			table
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string						false	"Table name"
//	@Param			table	body		database.TableJSONValues	false	"table values"
//	@Success		200		{object}	database.TableJSONValues
//	@Router			/databases/{name}/new_table [post]
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
//
//	@Summary		Add new row
//	@Description	Add new row
//	@Tags			table
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string		false	"Database name"
//	@Param			table	path		string		false	"Table name"
//	@Param			row		body		[]string	false	"Row to add"
//	@Success		200		{object}	database.TableJSONValues
//	@Router			/databases/{name}/{table}/new_row [post]
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
//
//	@Summary		Delete a row
//	@Description	Delete a row
//	@Tags			table
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	false	"Database name"
//	@Param			table	path		string	false	"Table name"
//	@Param			rowID	path		int		false	"row id"
//	@Success		200		{object}	database.TableJSONValues
//	@Router			/databases/{name}/{table}/{id} [delete]
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
//
//	@Summary		Delete a row
//	@Description	Delete a row
//	@Tags			table
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string		false	"Database name"
//	@Param			table	path		string		false	"Table name"
//	@Param			rowID	path		int			false	"row id"
//	@Param			row		body		[]string	false	"Row to add"
//	@Success		200		{object}	database.TableJSONValues
//	@Router			/databases/{name}/{table}/{id} [put]
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

// @Summary		Get databases list
// @Description	Get databases list
// @Tags			database
// @Accept			json
// @Produce		json
// @Success		200	{object}	[]string
// @Router			/databases [get]
func getDatabases(c echo.Context) error {
	response := database.ReadDatabasesPaths()
	return c.JSON(http.StatusOK, response)
}

// @Summary		Get tables list
// @Description	Get tables list
// @Tags			database
// @Accept			json
// @Produce		json
// @Param			name	path		string	false	"Database name"
// @Success		200		{object}	[]string
// @Router			/databases/{name} [get]
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

// @Summary		Get joined table
// @Description	Get table result of inner join two tables
// @Tags			table
// @Accept			json
// @Produce		json
// @Param			name	path		string	false	"Database name"
// @Param			t1		query		string	false	"First table"
// @Param			t2		query		string	false	"Second table"
// @Param			c1		query		string	false	"Column from first table"
// @Param			c2		query		string	false	"Column from second table"
// @Success		200		{object}	database.TableJSONValues
// @Router			/databases/{name}/joined_tables [get]
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

// @Summary		Delete database
// @Description	Delete database
// @Tags			database
// @Accept			json
// @Produce		json
// @Param			name	path		string	false	"Database name"
// @Success		200		{object}	string
// @Router			/databases/{name} [delete]
func deleteDB(c echo.Context) error {
	databaseName := c.Param("name")
	err := database.DeleteDatabase(databaseName)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, "deleted")
}

// @Summary		Delete table
// @Description	Delete table
// @Tags			table
// @Accept			json
// @Produce		json
// @Param			name	path		string	false	"Database name"
// @Param			table	path		string	false	"Table name"
// @Success		200		{object}	string
// @Router			/databases/{name}/{table} [delete]
func deleteTable(c echo.Context) error {
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

	err = db.DeleteTable(tableName)

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
