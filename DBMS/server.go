package main

import (
	"DBMS/database"
	"DBMS/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/fs"
	"log"
	"net/http"
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

	e.POST("/users", func(c echo.Context) error { return nil })
	e.GET("/users/:id", getUser)
	e.GET("/databases", getDatabases)
	e.GET("/databases/:name", getTables)
	e.GET("/databases/:name/:table", getTable)

	e.PUT("/users/:id", func(c echo.Context) error { return nil })
	e.DELETE("/users/:id", func(c echo.Context) error { return nil })

	e.Logger.Fatal(e.Start(":1323"))
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func getDatabases(c echo.Context) error {
	response := map[string]interface{}{
		"databases": []string{"aboba", "amogus", "EEEEE"},
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
		log.Println(err.Error())
	} else {
		log.Println("Saved")
	}

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
