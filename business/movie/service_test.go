package movie_test

import (
	"stock/business"
	"stock/business/movie"
	movieMock "stock/business/movie/mocks"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	title  = "Batman Beyond: Return of the Joker"
	year   = "2000"
	imdbID = "tt0233298"
	types  = "movie"
	poster = "https://m.media-amazon.com/images/M/MV5BNmRmODEwNzctYzU1MS00ZDQ1LWI2NWMtZWFkNTQwNDg1ZDFiXkEyXkFqcGdeQXVyNTI4MjkwNjA@._V1_SX300.jpg"
)

var (
	movieService    movie.Service
	movieRepository movieMock.Repository

	movieData    movie.MovieGeneral
	movieDataAll []movie.MovieGeneral
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

// func TestFindMovieByID(t *testing.T) {
// 	t.Run("Expect found the movie", func(t *testing.T) {
// 		movieRepository.On("FindMovieByID", mock.AnythingOfType("int")).Return(&movieData, nil).Once()

// 		movie, err := movieService.FindMovieByID(movieId)

// 		assert.Nil(t, err)
// 		assert.NotNil(t, movie)
// 		assert.Equal(t, name, movie.MovieName)

// 	})

// 	t.Run("Expect movie not found", func(t *testing.T) {
// 		movieRepository.On("FindMovieByID", mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()

// 		movie, err := movieService.FindMovieByID(movieId)

// 		assert.NotNil(t, err)

// 		assert.Nil(t, movie)

// 		assert.Equal(t, err, business.ErrNotFound)
// 	})

// }

func TestFindMovieByKeyword(t *testing.T) {
	t.Run("Expect found the movie", func(t *testing.T) {
		movieRepository.On("FindMovieByKeyword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&movieDataAll, nil).Once()

		movies, err := movieService.FindMovieByKeyword("batman", "2")

		assert.Nil(t, err)

		assert.NotNil(t, movies)

		assert.Equal(t, title, (*movies)[0].Title)
	})

	t.Run("Expect error Invalid spec", func(t *testing.T) {
		movieRepository.On("FindMovieByKeyword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrInvalidSpec).Once()

		movies, err := movieService.FindMovieByKeyword("", "")

		assert.Nil(t, movies)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInvalidSpec)
	})

	t.Run("Expect error Third party", func(t *testing.T) {
		movieRepository.On("FindMovieByKeyword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrThirdParty).Once()

		movies, err := movieService.FindMovieByKeyword("batman", "2")

		assert.Nil(t, movies)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrThirdParty)
	})
}

func setup() {

	movieData.Title = title
	movieData.ImdbID = imdbID
	movieData.Year = year
	movieData.Poster = poster
	movieData.Type = types

	movieDataAll = append(movieDataAll, movieData)

	movieService = movie.NewService(&movieRepository)
}
