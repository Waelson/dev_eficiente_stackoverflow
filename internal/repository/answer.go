package repository

import (
	"context"
	"database/sql"
	"github.com/Waelson/internal/model"
	"github.com/Waelson/internal/util/api"
)

type AnswerRepository interface {
	WithTransaction(ctx context.Context, fn TxFn) (err error)
	Save(ctx context.Context, answer *model.Answer) (*model.Answer, api.Error)
}

type answerRepository struct {
	database *sql.DB
	BaseRepository
}

func (a *answerRepository) Save(ctx context.Context, answer *model.Answer) (*model.Answer, api.Error) {
	sql := "INSERT INTO answers (post_id, response, user, create_at) VALUES (?,?,?,?)"

	statement, err := a.database.Prepare(sql)
	if err != nil {
		return nil, api.NewDatabaseError(err)
	}

	result, err := statement.Exec(answer.PostId, answer.Response, answer.User, answer.CreateAt)
	if err != nil {
		return nil, api.NewDatabaseError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, api.NewDatabaseError(err)
	}
	
	answer.Id = id
	return answer, nil
}

func NewAnswerRepository(database *sql.DB) AnswerRepository {
	return &answerRepository{
		database:       database,
		BaseRepository: NewBaseRepository(database),
	}
}
