package api

import (
	"stock/api/v1/movie"
	"sync"

	"github.com/labstack/echo"
)

func RegisterPath(e *echo.Echo, movieController *movie.Controller) {
	if movieController == nil {
		panic("Controller parameter cannot be nil")
	}
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)

	movieV1 := e.Group("v1/movies")

	go func() {
		movieV1.GET("", movieController.FindMovieByKeyword)
		waitgroup.Done()
	}()

	go func() {
		movieV1.GET("/:id", movieController.FindMovieByID)
		waitgroup.Done()
	}()

	waitgroup.Wait()

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
