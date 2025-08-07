package handler

import (
	statushttp "github.com/Hot-One/monolith/api/status_http"
	"github.com/Hot-One/monolith/config"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	cfg  *config.Config
	log  logger.Logger
	srvc service.ServiceInterface
}

func NewHandler(srvc service.ServiceInterface, config *config.Config, logger logger.Logger) *handler {
	return &handler{
		cfg:  config,
		log:  logger,
		srvc: srvc,
	}
}

func (h *handler) handleResponse(c *gin.Context, status statushttp.Status, data any) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
		)
	case code < 400:
		h.log.Warn(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
		)
	default:
		h.log.Error(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
		)
	}

	c.JSON(status.Code, statushttp.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}
