package movie

type Service interface {
	FindMovieByKeyword(keyword, page string) (*[]MovieGeneral, error)
	FindMovieByID(id string) (*MovieDetail, error)
}

type Repository interface {
	FindMovieByKeyword(keyword, page string) (*[]MovieGeneral, error)
	FindMovieByID(id string) (*MovieDetail, error)
}
