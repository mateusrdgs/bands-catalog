package controllers

import (
	"bands-catalog/models"
	"database/sql"

	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// GetAlbums is a controller to get all albums into catalog database
func GetAlbums(db *sql.DB) func(c echo.Context) error {

	return func(c echo.Context) error {
		rows, err := db.Query(`
		SELECT a.uuid, a.name, a.type, a.release_date, a.label, b.uuid as band_uuid
		FROM albums AS a
		INNER JOIN bands_albums AS ba
		ON a.uuid LIKE ba.album_uuid
		INNER JOIN bands AS b
		ON ba.band_uuid LIKE b.uuid
	`)

		if err != nil {
			panic(err.Error())
		}

		defer rows.Close()

		albums := make([]models.Album, 0)

		for rows.Next() {
			var album models.Album

			rows.Scan(&album.UUID, &album.Name, &album.Type, &album.ReleaseDate, &album.Label, &album.BandUUID)

			albums = append(albums, album)
		}

		return c.JSON(http.StatusOK, albums)

	}

}

// GetAlbum is a controller to get the specified album into catalog database
func GetAlbum(db *sql.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")

		if id == "" {
			message := "Album ID is missing"
			return c.JSON(http.StatusBadRequest, message)
		}

		query := `
			SELECT a.uuid, a.name, a.type, a.release_date, a.label, ba.band_uuid
			FROM albums as a
			INNER JOIN bands_albums as ba
			ON a.uuid LIKE ba.album_uuid
			WHERE a.uuid LIKE ?
		`
		var album models.Album

		row := db.QueryRow(query, id)

		err := row.Scan(&album.UUID, &album.Name, &album.Type, &album.ReleaseDate, &album.Label, &album.BandUUID)

		if err != nil {
			message := "No album was found"
			return c.JSON(http.StatusNotFound, message)
		}

		return c.JSON(http.StatusOK, album)
	}
}

// InsertAlbum is a controller to save incoming albums into catalog database
func InsertAlbum(db *sql.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		var album models.Album

		err := c.Bind(&album)

		if err != nil {
			panic(err.Error())
		}

		insert, err := db.Query(
			"CALL insert_album(?, ?, ?, ?, ?)",
			album.Name, album.Type, album.ReleaseDate, album.Label, album.BandUUID,
		)

		defer insert.Close()

		if err != nil {
			panic(err.Error())
		}

		message := strings.Join([]string{album.Name, "was saved successfully"}, " ")
		return c.JSON(http.StatusOK, message)

	}
}
