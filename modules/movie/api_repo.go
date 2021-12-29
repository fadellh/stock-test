package movie

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stock/business/movie"
)

type ApiRepository struct {
}

func NewApiRepo() ApiRepository {
	var ar ApiRepository
	return ar
}

// func (col *MovieGeneral) ToMovie() movie.Movie {
// 	var movie movie.Movie

// 	movie.Title = col.Title
// 	movie.Year = col.Year
// 	movie.ImdbID = col.ImdbID
// 	movie.Type = col.Type
// 	movie.Poster = col.Poster

// 	return movie
// }

func (ar ApiRepository) FindMovieByKeyword(keyword, page string) (*[]movie.MovieGeneral, error) {

	req, err := http.NewRequest("GET", "http://www.omdbapi.com/", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("apikey", "faf7e5bb")
	q.Add("s", keyword)
	q.Add("page", page)
	req.URL.RawQuery = q.Encode()

	response, err := http.Get(req.URL.String())

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var moviesRes movie.MovieSearchResponse
	json.Unmarshal(responseData, &moviesRes)

	return &moviesRes.Search, nil
}

func (ar ApiRepository) FindMovieByID(id string) (*movie.MovieDetail, error) {

	req, err := http.NewRequest("GET", "http://www.omdbapi.com/", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("apikey", "faf7e5bb")
	q.Add("i", id)
	req.URL.RawQuery = q.Encode()

	response, err := http.Get(req.URL.String())

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var moviesRes movie.MovieDetail
	json.Unmarshal(responseData, &moviesRes)

	return &moviesRes, nil
}
