package ledger

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

func (s *Service) SealDeed(config config.Configuration, deed entity.Deed) (entity.Deed, error) {
	return s.provider.SealDeed(config, deed)
}
