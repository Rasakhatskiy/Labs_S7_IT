package main

import (
	"DBMS/database"
	"DBMS/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		"databases": []string{"aboba", "amogus", "sus"},
	}
	return c.JSON(http.StatusOK, response)
}

func getTables(c echo.Context) error {
	databaseName := c.Param(":name")
	db, err := database.LoadDatabase(databaseName)
	tables := db.GetTablesList()

	if err != nil {
		log.Fatal("😭😭😭😭😭")
	}

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

	return c.JSON(http.StatusOK, table)
}
