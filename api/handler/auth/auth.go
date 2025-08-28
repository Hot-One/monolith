package auth_handler

import (
	statushttp "github.com/Hot-One/monolith/api/status_http"
	"github.com/Hot-One/monolith/config"
	auth_dto "github.com/Hot-One/monolith/dto/auth"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/service"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	cfg  *config.Config
	log  logger.Logger
	srvc service.ServiceInterface
}

func NewAuthHandler(group *gin.RouterGroup, cfg *config.Config, log logger.Logger, srvc service.ServiceInterface) {
	handler := &authHandler{
		cfg:  cfg,
		log:  log,
		srvc: srvc,
	}

	auth := group.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/logout", handler.Logout)
	}
}

// Login 			godoc
// @Summary      	Login
// @Description 	Login
// @Tags         	Auth Service
// @Accept       	json
// @Produce      	json
// @Param 			input body auth_dto.LoginRequest true "Login Request"
// @Success 		200 {object} auth_dto.LoginResponse "Login successful"
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router       	/auth/login [post]
func (h *authHandler) Login(c *gin.Context) {
	var input auth_dto.LoginRequest
	{
		if err := c.ShouldBindJSON(&input); err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	token, err := h.srvc.AuthService().Login(c.Request.Context(), &input)
	{
		if err != nil {
			statushttp.InternalServerError(c, err.Error())
			return
		}
	}

	statushttp.OK(c, token)
}

func (h *authHandler) Logout(c *gin.Context) {

}
