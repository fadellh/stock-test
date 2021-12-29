package movie

import (
	"stock/business"
)

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo,
	}
}

func (s *service) FindMovieByKeyword(keyword, page string) (*[]MovieGeneral, error) {

	if keyword == "" {
		return nil, business.ErrInvalidSpec
	}

	if page == "" {
		page = "1"
	}

	res, err := s.repository.FindMovieByKeyword(keyword, page)

	if err != nil {
		return nil, business.ErrThirdParty
	}

	return res, nil

}

func (s *service) FindMovieByID(id string) (*MovieDetail, error) {

	if id == "" {
		return nil, business.ErrInvalidSpec
	}

	res, err := s.repository.FindMovieByID(id)

	if res == nil {
		return nil, business.ErrNotFound
	}

	if err != nil {
		return nil, business.ErrThirdParty
	}

	return res, nil

}
