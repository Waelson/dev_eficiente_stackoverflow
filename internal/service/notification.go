package service

import (
	"context"
	"github.com/Waelson/internal/model"
)

type NotificationService interface {
	Notify(ctx context.Context, answer *model.Answer) error
}

type notificationService struct {
}

func (s *notificationService) Notify(ctx context.Context, answer *model.Answer) error {
	return nil
}

func NewNotificationService() NotificationService {
	return &notificationService{}
}
