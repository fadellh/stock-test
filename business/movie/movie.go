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

type MovieDetail struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	ImdbID   string `json:"imdbID"`
	Type     string `json:"Type"`
	Director string `json:"Director"`
	Plot     string `json:"Plot"`
	Genre    string `json:"Genre"`
	Runtime  string `json:"Runtime"`
	Awards   string `json:"Awards"`
	Ratings  []Ratings
}

type Ratings struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}
