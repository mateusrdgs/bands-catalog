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

		bands := make([]models.Band, 0)

		for rows.Next() {
			var band models.Band

			rows.Scan(&band.UUID, &band.Name, &band.YearOfFoundation, &band.Biography, &band.Country, &band.Genre)

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

// GetBand is a controller to fetch a specific band into the database
func GetBand(db *sql.DB) func(c echo.Context) error {

	return func(c echo.Context) error {

		id := c.Param("id")

		if id == "" {
			message := "Id wasn't informed"
			return c.JSON(http.StatusBadRequest, message)
		}

		query := "SELECT * FROM bands WHERE uuid LIKE ?"

		var band models.Band
		err := db.QueryRow(query, id).Scan(&band.UUID, &band.Name, &band.YearOfFoundation, &band.Biography, &band.Country, &band.Genre)

		if err != nil {
			return c.JSON(http.StatusNotFound, "Band not found!")
		}

		return c.JSON(http.StatusOK, band)

	}

}
