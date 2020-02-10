package controllers

import (
	"bands-catalog/models"
	"database/sql"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// GetBands is a controller to save incoming bands into catalog database
func GetBands(db *sql.DB) func(c echo.Context) error {

	return func(c echo.Context) error {
		rows, err := db.Query("SELECT * FROM bands")

		if err != nil {
			panic(err.Error())
		}

		defer rows.Close()

		bands := make([]models.Band, 0, 6)

		for rows.Next() {
			var uuid, name, biography, country, genre string
			var yearOfFoundation int
			var band models.Band

			rows.Scan(&uuid, &name, &yearOfFoundation, &biography, &country, &genre)

			band.Name = name
			band.YearOfFoundation = yearOfFoundation
			band.Biography = biography
			band.Country = country
			band.Genre = genre

			bands = append(bands, band)
		}

		return c.JSON(http.StatusOK, bands)
	}
}

// InsertBand is a controller to save incoming bands into catalog database
func InsertBand(db *sql.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		band := new(models.Band)
		err := c.Bind(band)

		if err != nil {
			panic(err.Error())
		}

		if band.Name == "" {
			message := "Missing parameter 'name'"
			return c.String(http.StatusBadRequest, message)
		}

		if band.YearOfFoundation == 0 {
			message := "Missing parameter 'year_of_foundation'"
			return c.String(http.StatusBadRequest, message)
		}

		if band.Biography == "" {
			message := "Missing parameter 'biography'"
			return c.String(http.StatusBadRequest, message)
		}

		if band.Country == "" {
			message := "Missing parameter 'country'"
			return c.String(http.StatusBadRequest, message)
		}

		if band.Genre == "" {
			message := "Missing parameter 'genre'"
			return c.String(http.StatusBadRequest, message)
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
	}
}
