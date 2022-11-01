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

	e.GET("/databases", getDatabases)
	e.GET("/databases/:name", getTables)
	e.GET("/databases/:name/:table", getTable)

	e.POST("/databases/:name/:table/new_row", addRow)
	e.DELETE("/databases/:name/:table/:rowID", deleteRow)
	e.PUT("/databases/:name/:table/:rowID", editRow)

	e.DELETE("/users/:id", func(c echo.Context) error { return nil })

	e.Logger.Fatal(e.Start(":1323"))

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

//e.PUT("/databases/:name/:table/:rowID", editRow)
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
	response := map[string]interface{}{
		"databases": []string{"cats", "dogs", "EEEEE"},
	}
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

	var headers []JsonTableHeader
	for i, _ := range table.Headers {
		headers = append(headers, JsonTableHeader{
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

	jt := JsonTable{
		Headers: headers,
		Values:  interValues,
	}

	return c.JSON(http.StatusOK, jt)
}

type JsonTableHeader struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type JsonTable struct {
	Headers []JsonTableHeader `json:"headers"`
	Values  [][]interface{}   `json:"values"`
}

//type JsonTable struct {
//	Headers struct {
//		Name []string `json:"name"`
//		Type []string `json:"type"`
//	} `json:"headers"`
//	Values [][]database.DBType `json:"values"`
//}
