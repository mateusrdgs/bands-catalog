package routes

import (
	"bands-catalog/controllers"
	"bands-catalog/models"
	"database/sql"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// StartAlbumsRouting is a routing handler to albums context inside of routes package
func StartAlbumsRouting(e *echo.Echo, db *sql.DB) {
	e.GET("/albums", controllers.GetAlbums(db))

	g := e.Group("/album")
	g.POST("/", controllers.InsertAlbum(db))
	g.GET("/:id", controllers.GetAlbum(db))

	g.POST("/:album_uuid/song", func(c echo.Context) error {
		song := new(models.Song)
		albumID := c.Param("album_uuid")
		err := c.Bind(song)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Some error occurred, try again later")
		}

		query := "CALL insert_song (?, ?, ?, ?, ?)"

		_, dbErr := db.Exec(query, song.Name, song.Number, song.Duration, song.Lyrics, albumID)

		if dbErr != nil {
			return c.JSON(http.StatusInternalServerError, "Some error occurred, try again later")
		}

		message := strings.Join([]string{song.Name, "was saved successfully"}, " ")
		return c.JSON(http.StatusOK, message)

	})

	g.GET("/:album_uuid/songs", func(c echo.Context) error {
		albumID := c.Param("album_uuid")
		songs := make([]models.Song, 0)

		query := `
			SELECT DISTINCT s.uuid, s.name, s.number, s.duration, s.lyrics
			FROM albums AS a
			INNER JOIN albums_songs AS a_s
			ON ? LIKE a_s.album_uuid
			INNER JOIN songs AS s
			ON a_s.song_uuid LIKE s.uuid
			ORDER BY s.number;
		`

		rows, err := db.Query(query, albumID)

		defer rows.Close()

		if err != nil {
			message := "Some error occurred, try again later"
			c.JSON(http.StatusInternalServerError, message)
		}

		for rows.Next() {
			var song models.Song
			rows.Scan(&song.UUID, &song.Name, &song.Number, &song.Duration, &song.Lyrics)
			songs = append(songs, song)
		}

		return c.JSON(http.StatusOK, songs)

	})

}
