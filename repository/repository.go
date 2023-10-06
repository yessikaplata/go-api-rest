package repository

import (
	"database/sql"

	"github.com/yessikaplata/go-api-rest/db"
	"github.com/yessikaplata/go-api-rest/model"
)

type Repository struct {
	Db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		Db: db,
	}
}

func (r Repository) Create(user model.User) (model.User, error) {
	sql := "INSERT users SET user_name=?, password=?, email=?"
	result, err := r.Db.Exec(sql, user.UserName, user.Password, user.Email)
	if err != nil {
		return model.User{}, db.HandleError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return model.User{}, db.HandleError(err)
	}
	user.Id = int(id)
	return user, nil
}

func (r Repository) Find(id int) (model.User, error) {
	var user model.User
	query := "SELECT id, user_name,password,email FROM users where id = ?"
	rows, err := r.Db.Query(query, id)
	if err != nil {
		return model.User{}, db.HandleError(err)
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email); err != nil {
			return model.User{}, db.HandleError(err)
		}
		return user, nil
	}
	return user, db.ErrObjectNotFound{}
}

func (r Repository) Update(user model.User) (model.User, error) {
	sql := "UPDATE users SET user_name=?, password=?, email=? WHERE id=?"
	_, err := r.Db.Exec(sql, user.UserName, user.Password, user.Email, user.Id)

	if err != nil {
		return model.User{}, db.HandleError(err)
	}
	return user, nil
}

func (r Repository) Delete(id int) error {
	query := "DELETE FROM users where id = ?"
	result, err := r.Db.Exec(query, id)
	if err != nil {
		return db.HandleError(err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return db.ErrObjectNotFound{}
	}
	return nil
}

func (r Repository) FindAll() ([]model.User, error) {
	query := "SELECT id, user_name,password,email FROM users"
	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, db.HandleError(err)
	}
	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email); err != nil {
			return nil, db.HandleError(err)
		}
		users = append(users, user)
	}
	return users, nil
}
