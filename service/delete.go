package service

import (
	"context"
	"errors"

	"github.com/yessikaplata/go-api-rest/db"
	service_errors "github.com/yessikaplata/go-api-rest/service/errors"
)

func (s Service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return service_errors.NewUserNotFoundError()
	default:
		return err
	}
	return nil
}
