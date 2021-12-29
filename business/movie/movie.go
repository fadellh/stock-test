package movie

type Movie struct {
	Title  string
	Year   string
	ImdbID string
	Type   string
	Poster string
}

type MovieSearchResponse struct {
	Search []MovieGeneral `json:"Search"`
}

type MovieGeneral struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}
