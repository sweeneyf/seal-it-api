package ledger

import (
	"github.com/sweeneyf/seal-it-api/entity"
	"github.com/sweeneyf/seal-it-api/pkg/config"
)

type UseCase interface {
	SealDeed(config.Configuration, entity.Deed) (entity.Deed, error)
}
