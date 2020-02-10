package routes

import (
	"bands-catalog/models"
	"database/sql"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// StartBandsRouting is a function to start the routing of bands module
func StartBandsRouting(e *echo.Echo, db *sql.DB) {

	e.POST("/band", func(c echo.Context) error {

		band := new(models.Band)
		err := c.Bind(band)

		if err != nil {
			panic(err.Error())
		}

		insert, err := db.Query(
			"CALL insert_band(?, ?, ?, ?, ?)",
			band.Name, band.YearOfFoundation, band.Biography, band.Country, band.Genre,
		)

		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()

		message := strings.Join([]string{band.Name, "was saved successfully"}, " ")

		return c.String(http.StatusOK, message)

	})

}
