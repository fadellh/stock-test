package api

import (
	"stock/api/v1/movie"

	"github.com/labstack/echo"
)

func RegisterPath(e *echo.Echo, movieController *movie.Controller) {
	if movieController == nil {
		panic("Controller parameter cannot be nil")
	}

	movieV1 := e.Group("v1/movies")
	movieV1.GET("", movieController.FindMovieByKeyword)
	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
