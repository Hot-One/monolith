package user_handler

import (
	statushttp "github.com/Hot-One/monolith/api/status_http"
	"github.com/Hot-One/monolith/config"
	user_dto "github.com/Hot-One/monolith/dto/user"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/pkg/pg"
	"github.com/Hot-One/monolith/service"
	user_service "github.com/Hot-One/monolith/service/user"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	cfg  *config.Config
	log  logger.Logger
	srvc user_service.UserServiceInterface
}

func NewUserHandler(group *gin.RouterGroup, srvc service.ServiceInterface, config *config.Config, logger logger.Logger) {
	handler := &userHandler{
		cfg:  config,
		log:  logger,
		srvc: srvc.UserService(),
	}

	user := group.Group("/user")
	{
		user.POST("/", handler.Create)
		// user.GET("/:id", handler.GetByID)
		// user.PUT("/:id", handler.Update)
		// user.DELETE("/:id", handler.Delete)
	}
}

// Create 			godoc
// @Summary 		user-service-create
// @Description 	user-service-create
// @Tags 			User Service
// @Accept 			json
// @Produce 		json
// @Param 			user body user_dto.UserCreate true "User Create"
// @Success 		201 {object} statushttp.Response "User created successfully"
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router 			/user [post]
func (h *userHandler) Create(c *gin.Context) {
	var in user_dto.UserCreate

	{
		if err := c.ShouldBindJSON(&in); err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	id, err := h.srvc.Create(c.Request.Context(), &in)
	{
		if err != nil {
			statushttp.InternalServerError(c, err.Error())
			return
		}
	}

	statushttp.Created(c, pg.Id{Id: id})
}
