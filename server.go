package main

import (
	"database/sql"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/bands_catalog")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.POST("/band", func(c echo.Context) error {

		bandName := c.FormValue("band_name")
		yearOfFoundation := c.FormValue("year_of_foundation")
		biography := c.FormValue("biography")
		country := c.FormValue("country")
		genre := c.FormValue("genre")

		insert, err := db.Query("CALL save_band(?, ?, ?, ?, ?)", bandName, yearOfFoundation, biography, country, genre)

		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()

		message := strings.Join([]string{bandName, "was saved successfully"}, " ")

		return c.String(http.StatusOK, message)

	})

	e.Logger.Fatal(e.Start(":1234"))
}
