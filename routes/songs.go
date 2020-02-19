package routes

import (
	"bands-catalog/models"
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

// StartSongsRouting is a function to start the routing of songs module
func StartSongsRouting(e *echo.Echo, db *sql.DB) {
	e.GET("/songs/:song_uuid", func(c echo.Context) error {
		songID := c.Param("song_uuid")
		var song models.Song

		query := "SELECT * FROM songs WHERE uuid LIKE ?"

		row := db.QueryRow(query, songID)
		err := row.Scan(&song.UUID, &song.Name, &song.Number, &song.Duration, &song.Lyrics)

		if err != nil {
			message := "Some error occurred, try again later"
			return c.JSON(http.StatusInternalServerError, message)
		}

		return c.JSON(http.StatusOK, song)

	})
}
