package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	//createTestJson()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.POST("/users", func(c echo.Context) error { return nil })
	e.GET("/users/:id", getUser)
	e.GET("/databases", getDatabases)
	e.GET("/databases/:name/tables", getTables)

	e.PUT("/users/:id", func(c echo.Context) error { return nil })
	e.DELETE("/users/:id", func(c echo.Context) error { return nil })

	e.Logger.Fatal(e.Start(":1323"))
}

//type Entry struct {
//	Name string `json:"name"`
//	Path string `json:"path"`
//}
//
//func createTestJson() {
//	datas := make(map[string]Entry)
//	datas["gorillaz"] = Entry{Name: "gorillaz", Path: "databases/gorillaz"}
//	datas["diamonds"] = Entry{Name: "diamonds", Path: "databases/diamonds"}
//	datas["ff"] = Entry{Name: "ff", Path: "databases/ff"}
//
//	jsonString, err := json.Marshal(datas)
//	if err == nil {
//		_ = ioutil.WriteFile("databases.json", jsonString, 0644)
//	}
//}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func getDatabases(c echo.Context) error {
	response := map[string]interface{}{
		"databases": []string{"aboba", "amosus", "sus"},
	}
	return c.JSON(http.StatusOK, response)
}

func getTables(c echo.Context) error {
	//databaseName := c.Param(":name")
	response := map[string]interface{}{
		"databases": []string{"aboba", "amosus", "sus"},
	}
	return c.JSON(http.StatusOK, response)
}
