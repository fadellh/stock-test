package movie

import (
	"stock/api/common"
	"stock/business/movie"

	"github.com/labstack/echo"
)

type Controller struct {
	service movie.Service
}

func NewController(service movie.Service) *Controller {
	return &Controller{
		service,
	}
}

func (controller *Controller) FindMovieByKeyword(c echo.Context) error {
	page := c.QueryParam("page")
	keyword := c.QueryParam("keyword")

	movies, err := controller.service.FindMovieByKeyword(keyword, page)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponse(movies))

}

func (controller *Controller) FindMovieByID(c echo.Context) error {

	id := (c.Param("id"))

	movie, err := controller.service.FindMovieByID(id)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponse(movie))
}
