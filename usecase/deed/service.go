package deed

import (
	"github.com/sweeneyf/seal-it-api/entity"
	"github.com/sweeneyf/seal-it-api/pkg/config"
)

type Service struct {
	provider UseCase
}

//NewService create new service
func NewService(provider UseCase) *Service {
	return &Service{
		provider: provider,
	}
}

func (s *Service) SaveDeed(config config.Configuration, deed entity.Deed) error {
	return s.provider.SaveDeed(config, deed)
}
