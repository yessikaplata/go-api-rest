package service

import (
	"context"

	"github.com/asaskevich/govalidator"
	"github.com/yessikaplata/go-api-rest/model"
	"github.com/yessikaplata/go-api-rest/service/errors"
)

type CreateParams struct {
	UserName string `valid:"required"`
	Password string `valid:"required"`
	Email    string `valid:"required"`
}

func (s Service) Create(ctx context.Context, params CreateParams) (model.User, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return model.User{}, errors.NewArgumentsError()
	}

	tx, err := s.repository.Db.BeginTx(ctx, nil)
	if err != nil {
		return model.User{}, err
	}
	defer tx.Rollback()

	user := model.User{
		UserName: params.UserName,
		Password: params.Password,
		Email:    params.Email,
	}
	user, err = s.repository.Create(user)
	if err != nil {
		return model.User{}, err
	}

	err = tx.Commit()
	return user, err
}
