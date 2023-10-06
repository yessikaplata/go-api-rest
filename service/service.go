package service

import "github.com/yessikaplata/go-api-rest/repository"

type Service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return Service{
		repository: repository,
	}
}
