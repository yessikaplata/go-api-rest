package service

import (
	"context"
	"errors"

	"github.com/yessikaplata/go-api-rest/db"
	"github.com/yessikaplata/go-api-rest/model"
	service_errors "github.com/yessikaplata/go-api-rest/service/errors"
)

func (s Service) Get(ctx context.Context, id int) (model.User, error) {
	user, err := s.repository.Find(id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return model.User{}, service_errors.NewUserNotFoundError()
	default:
		return model.User{}, err
	}
	return user, nil
}

func (s Service) GetAll(ctx context.Context) ([]model.User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
