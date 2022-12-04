package main

import (
	_ "DBMS/docs"
	"DBMS/genImpl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

//	@title			My database management system
//	@version		1.0
//	@description	🙀

//	@contact.name	Максим РАсахацький ТТП-42
//	@contact.email	saharok.maks@gmail.com

// @host		localhost:1323
// @accept		json
// @produce	json

func main()  {
	genImpl.SetupHandler()
}

func mainOld() {
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

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// DATABASES
	e.GET("/databases", getDatabases)
	e.POST("/databases/new_database", createDB)
	e.DELETE("/databases/:name", deleteDB)

	// TABLES
	e.GET("/databases/:name", getTables)
	e.GET("/databases/:name/:table", getTable)
	e.POST("/databases/:name/new_table", addTable)
	e.GET("/databases/:name/join_tables", getJoinTablesData)
	e.GET("/databases/:name/joined_tables", getJoinedTables)
	e.DELETE("/databases/:name/:table", deleteTable)

	// ROW
	e.POST("/databases/:name/:table/new_row", addRow)
	e.PUT("/databases/:name/:table/:rowID", editRow)
	e.DELETE("/databases/:name/:table/:rowID", deleteRow)

	e.Logger.Fatal(e.Start(":1323"))
}
