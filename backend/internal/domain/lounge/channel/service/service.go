package service

import "github.com/shjk0531/moye/backend/internal/domain/lounge/channel/repository"


type Service interface {
	GetChannelService() ChannelService
}

type service struct {
	repo repository.ChannelRepository
	channelService ChannelService
}

func NewService(repo repository.ChannelRepository) Service {
	return &service{
		repo: repo,
	}
}

func Init(repo repository.ChannelRepository) Service {
	svc := &service{
		repo: repo,
	}
	svc.channelService = NewChannelService(repo)
	return svc
}

func (s *service) GetChannelService() ChannelService {
	return s.channelService
}

