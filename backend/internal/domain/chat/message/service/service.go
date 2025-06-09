package service

import "github.com/shjk0531/moye/backend/internal/domain/chat/message/repository"

type Service interface {
	GetMessageService() MessageService
}

type service struct {
	repo repository.Repository
	messageService MessageService
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func Init(repo repository.Repository) Service {
	svc := &service{repo: repo}
	svc.messageService = NewMessageService(repo)
	return svc
}

func (s *service) GetMessageService() MessageService {
	return s.messageService
}