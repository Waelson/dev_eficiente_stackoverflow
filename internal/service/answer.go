package service

import (
	"context"
	"github.com/Waelson/internal/model"
	"github.com/Waelson/internal/repository"
	"github.com/Waelson/internal/util/api"
	"time"
)

const invalidStatusError = "invalid status"

type AnswerService interface {
	Save(ctx context.Context, answer *model.Answer) (*model.Answer, api.Error)
}

type answerService struct {
	answerRepository    repository.AnswerRepository
	notificationService NotificationService
	postRepository      repository.PostRepository
}

func (s *answerService) Save(ctx context.Context, answer *model.Answer) (*model.Answer, api.Error) {
	answer.CreateAt = time.Now()

	post, err := s.postRepository.GetById(ctx, answer.PostId)
	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, api.NewServiceError("post not found")
	}

	if post.IsClosed() {
		return nil, api.NewServiceError(invalidStatusError)
	}

	answer, err = s.answerRepository.Save(ctx, answer)
	if err != nil {
		return answer, err
	}

	go func() {
		s.notificationService.Notify(ctx, answer)
	}()

	return answer, nil
}

func NewAnswerService(answerRepository repository.AnswerRepository, postRepository repository.PostRepository, notificationService NotificationService) AnswerService {
	return &answerService{
		answerRepository:    answerRepository,
		postRepository:      postRepository,
		notificationService: notificationService,
	}
}
