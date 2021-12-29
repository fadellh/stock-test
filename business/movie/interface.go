package movie

type Service interface {
	FindMovieByKeyword(keyword, page string) (*[]MovieGeneral, error)
}

type Repository interface {
	FindMovieByKeyword(keyword, page string) (*[]MovieGeneral, error)
}
