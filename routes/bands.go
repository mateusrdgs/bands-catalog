package routes

import (
	"bands-catalog/controllers"
	"database/sql"

	"github.com/labstack/echo"
)

// StartBandsRouting is a function to start the routing of bands module
func StartBandsRouting(e *echo.Echo, db *sql.DB) {

	e.GET("/bands", controllers.GetBands(db))

	g := e.Group("/band")

	g.GET("/:id", controllers.GetBand(db))
	g.POST("/", controllers.InsertBand(db))

}
