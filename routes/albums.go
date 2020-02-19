package routes

import (
	"bands-catalog/controllers"
	"database/sql"

	"github.com/labstack/echo"
)

// StartAlbumsRouting is a routing handler to albums context inside of routes package
func StartAlbumsRouting(e *echo.Echo, db *sql.DB) {

	e.GET("/albums", controllers.GetAlbum(db))
	e.POST("/album", controllers.InsertAlbum(db))

}
