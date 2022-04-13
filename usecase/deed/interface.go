package deed

import (
	"github.com/sweeneyf/seal-it-api/entity"
	"github.com/sweeneyf/seal-it-api/pkg/config"
)

type UseCase interface {
	SaveDeed(config.Configuration, entity.Deed) error
}
