package repository

import (
	"context"
	"database/sql"
	"github.com/Waelson/internal/model"
	"github.com/Waelson/internal/util/api"
	"strings"
)

const TagSeparator = ","

type PostRepository interface {
	WithTransaction(ctx context.Context, fn TxFn) (err error)
	Save(ctx context.Context, post *model.Post) (*model.Post, api.Error)
	GetById(ctx context.Context, id int64) (*model.Post, api.Error)
}

type postRepository struct {
	database *sql.DB
	BaseRepository
}

func (a *postRepository) GetById(ctx context.Context, id int64) (*model.Post, api.Error) {
	query := "SELECT id, title, description, user, tags, create_at, status FROM posts WHERE id = ?"
	var post model.Post
	var tags string
	err := a.database.QueryRow(query, id).Scan(&post.Id, &post.Title, &post.Description, &post.User, &tags, &post.CreateAt, &post.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, api.NewDatabaseError(err)
	}
	post.Tags = strings.Split(tags, TagSeparator)
	return &post, nil
}

func (a *postRepository) Save(ctx context.Context, post *model.Post) (*model.Post, api.Error) {
	sql := "INSERT INTO posts (title, description, tags, user, create_at, status) VALUES (?,?,?,?,?,?)"
	statement, err := a.database.Prepare(sql)
	if err != nil {
		return nil, api.NewDatabaseError(err)
	}

	_, err = statement.Exec(post.Title, post.Description, strings.Join(post.Tags, TagSeparator), post.User, post.CreateAt, post.Status)
	if err != nil {
		return nil, api.NewDatabaseError(err)
	}
	return post, nil
}

func NewPostRepository(database *sql.DB) PostRepository {
	return &postRepository{
		database:       database,
		BaseRepository: NewBaseRepository(database),
	}
}
