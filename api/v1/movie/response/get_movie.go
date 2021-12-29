package response

import "stock/business/movie"

type GetMovieResponse struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

func NewGetMovieResponse(movie movie.Movie) *GetMovieResponse {
	var getMovieResponse GetMovieResponse

	getMovieResponse.Title = movie.Title
	getMovieResponse.Year = movie.Year
	getMovieResponse.ImdbID = movie.ImdbID
	getMovieResponse.Type = movie.Type
	getMovieResponse.Poster = movie.Poster

	return &getMovieResponse
}
