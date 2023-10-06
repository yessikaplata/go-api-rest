package service

import (
	"context"

	"github.com/asaskevich/govalidator"
	"github.com/yessikaplata/go-api-rest/model"
	"github.com/yessikaplata/go-api-rest/service/errors"
)

type UpdateParams struct {
	Id       int    `valid:"required"`
	UserName string `valid:"required"`
	Password string `valid:"required"`
	Email    string `valid:"required"`
}

func (s Service) Update(ctx context.Context, params UpdateParams) (model.User, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return model.User{}, errors.NewArgumentsError()
	}

	tx, err := s.repository.Db.BeginTx(ctx, nil)
	if err != nil {
		return model.User{}, err
	}
	defer tx.Rollback()

	userDB, err := s.repository.Find(params.Id)
	if userDB.Id == 0 {
		return model.User{}, errors.NewUserNotFoundError()
	}

	user := model.User{
		Id:       params.Id,
		UserName: params.UserName,
		Password: params.Password,
		Email:    params.Email,
	}
	user, err = s.repository.Update(user)
	if err != nil {
		return model.User{}, err
	}

	err = tx.Commit()
	return user, err
}
