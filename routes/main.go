package routes

import (
	"database/sql"

	"github.com/labstack/echo"
)

// StartRouting is a function to start the application routing
func StartRouting(e *echo.Echo, db *sql.DB) {

	StartBandsRouting(e, db)
	StartAlbumsRouting(e, db)
	StartSongsRouting(e, db)

}
