package repository

import (
	"context"
	"database/sql"
	"github.com/Waelson/internal/model"
	"github.com/Waelson/internal/util/api"
)

type UserRepository interface {
	Save(ctx context.Context, user *model.User) (*model.User, api.Error)
	FindByLogin(ctx context.Context, login string) (*model.User, api.Error)
	FindByEmail(ctx context.Context, email string) (*model.User, api.Error)
	ListByLogins(ctx context.Context, logins []string) (*[]model.User, api.Error)
}

type userRepository struct {
	database *sql.DB
	BaseRepository
}

func (r *userRepository) Save(ctx context.Context, user *model.User) (*model.User, api.Error) {
	sql := "INSERT INTO users (name, login, email, password, create_at) VALUES (?,?,?,?,?)"

	statement, err := r.database.Prepare(sql)
	if err != nil {
		return nil, api.NewDatabaseError(err)
	}

	result, err := statement.Exec(user.Name, user.Login, user.Email, user.Password, user.CreateAt)
	if err != nil {
		return nil, api.NewDatabaseError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, api.NewDatabaseError(err)
	}

	user.Id = id
	return user, nil
}

func (r *userRepository) FindByLogin(ctx context.Context, login string) (*model.User, api.Error) {
	query := "SELECT id, name, login, email, password, create_at FROM users WHERE login = ?"
	var user model.User
	err := r.database.QueryRow(query, login).Scan(&user.Id, &user.Name, &user.Login, &user.Email, &user.Password, &user.CreateAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, api.NewDatabaseError(err)
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, api.Error) {
	query := "SELECT id, name, login, email, password, create_at FROM users WHERE email = ?"
	var user model.User
	err := r.database.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Login, &user.Email, &user.Password, &user.CreateAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, api.NewDatabaseError(err)
	}
	return &user, nil
}

func (r *userRepository) ListByLogins(ctx context.Context, logins []string) (*[]model.User, api.Error) {
	return nil, nil
}

func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepository{
		database:       database,
		BaseRepository: NewBaseRepository(database),
	}
}
