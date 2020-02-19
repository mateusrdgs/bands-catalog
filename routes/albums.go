package routes

import (
	"bands-catalog/controllers"
	"database/sql"

	"github.com/labstack/echo"
)

// StartAlbumsRouting is a routing handler to albums context inside of routes package
func StartAlbumsRouting(e *echo.Echo, db *sql.DB) {
	e.GET("/albums", controllers.GetAlbums(db))

	g := e.Group("/album")

	g.GET("/:id", controllers.GetAlbum(db))
	g.POST("/", controllers.InsertAlbum(db))
}
