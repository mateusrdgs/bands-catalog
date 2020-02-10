package main

import (
	"bands-catalog/database"
	"bands-catalog/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	db, err := database.StartDatabase("mysql", "root@tcp(127.0.0.1:3306)/bands_catalog")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	routes.StartRouting(e, db)

	e.Logger.Fatal(e.Start(":1234"))
}
