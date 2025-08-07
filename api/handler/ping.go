package handler

import (
	statushttp "github.com/Hot-One/monolith/api/status_http"
	"github.com/gin-gonic/gin"
)

// Ping godoc
// @ID 					ping
// @Router 				/v1/ping [get]
// @Tags 				Service
// @Accept  			json
// @Produce  			json
// @Summary 			Ping service to check if it's running
// @Description 		Ping service to check if it's running
// @Success 200			{object} statushttp.Response "Service is running"
// @Response 400 		{object} statushttp.Response "Bad Request"
// @Failure 500 		{object} statushttp.Response "Internal Server Error"
func (h *handler) Ping(c *gin.Context) {
	h.handleResponse(c, statushttp.OK, "pong")
}
