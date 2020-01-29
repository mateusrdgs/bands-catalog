package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/bands_catalog")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// insert, err := db.Query("CALL save_band('Craft', 1998, 'Craft did not perform live until September 2014.', 'Sweden', 'Black Metal')")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()

	e.Logger.Fatal(e.Start(":1234"))
}
