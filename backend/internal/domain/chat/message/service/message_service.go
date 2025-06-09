package service

import (
	"github.com/shjk0531/moye/backend/internal/domain/chat/message/model"
	"github.com/shjk0531/moye/backend/internal/domain/chat/message/repository"
)

type MessageService interface {
	SaveMessage(channelID string, userID string, content string, msgType string) (*model.Message, error)
	FindMessagesByChatRoomID(chatRoomID int) ([]*model.Message, error)
}

type messageService struct {
	repo repository.Repository
}

func NewMessageService(repo repository.Repository) MessageService {
	return &messageService{repo: repo}
}

func (s *messageService) SaveMessage(channelID string, userID string, content string, msgType string) (*model.Message, error) {
	message := &model.Message{
		ChannelID: channelID,
		UserID:    userID,
		Content:   content,
		Type:      msgType,
	}

	err := s.repo.InsertMessage(message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (s *messageService) FindMessagesByChatRoomID(chatRoomID int) ([]*model.Message, error) {
	return s.repo.FindMessagesByChatRoomID(uint(chatRoomID))
}