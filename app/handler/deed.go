package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sweeneyf/seal-it-api/app/presenter"
	"github.com/sweeneyf/seal-it-api/entity"
	"github.com/sweeneyf/seal-it-api/pkg/config"
	"github.com/sweeneyf/seal-it-api/pkg/logger"
	"github.com/sweeneyf/seal-it-api/usecase/deed"
	"github.com/sweeneyf/seal-it-api/usecase/ledger"
)

type DeedHandler struct {
	deedService   deed.UseCase
	ledgerService ledger.UseCase
	config        config.Configuration
}

func NewDeedHandler(ledgerService ledger.UseCase, deedService deed.UseCase, config config.Configuration) *DeedHandler {
	return &DeedHandler{
		config:        config,
		deedService:   deedService,
		ledgerService: ledgerService,
	}
}

func (h *DeedHandler) SealAndSaveDeed(c *gin.Context) {
	var newDeedRequest presenter.DeedRequest

	if err := c.BindJSON(&newDeedRequest); err != nil {
		apiErr := presenter.NewBadRequestError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
	}
	deed := entity.Deed{}
	logger.Log.Info("posting a new deed")
	deed, err := h.ledgerService.SealDeed(h.config, deed)
	if err != nil {

	}
	h.deedService.SaveDeed(h.config, deed)
	c.IndentedJSON(http.StatusCreated, deed)

}
