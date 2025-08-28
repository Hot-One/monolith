package user_handler

import (
	statushttp "github.com/Hot-One/monolith/api/status_http"
	"github.com/Hot-One/monolith/config"
	user_dto "github.com/Hot-One/monolith/dto/user"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/pkg/pg"
	"github.com/Hot-One/monolith/pkg/utils"
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
		user.POST("", handler.Create)
		user.PATCH("/:id", handler.Update)
		user.GET("", handler.GetList)
		user.GET("/:id", handler.GetById)
		user.GET("/search", handler.Search)
		user.DELETE("/:id", handler.Delete)
	}
}

// Create 			godoc
// @Summary 		user-service-create
// @Description 	user-service-create
// @Tags 			User Service
// @Accept 			json
// @Produce 		json
// @Param 			input body user_dto.UserCreate true "User Create"
// @Success 		201 {object} pg.Id "User created successfully"
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router 			/user [post]
func (h *userHandler) Create(c *gin.Context) {
	var input user_dto.UserCreate

	{
		if err := c.ShouldBindJSON(&input); err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	{
		if !utils.IsValidLogin(input.Username) {
			statushttp.BadRequest(c, "invalid username")
			return
		}

		if !utils.IsValidPassword(input.Password) {
			statushttp.BadRequest(c, "invalid password")
			return
		}
	}

	id, err := h.srvc.Create(c.Request.Context(), &input)
	{
		if err != nil {
			statushttp.InternalServerError(c, err.Error())
			return
		}
	}

	statushttp.Created(c, pg.Id{Id: id})
}

// Update 			godoc
// @Summary 		user-service-update
// @Description 	user-service-update
// @Tags 			User Service
// @Accept 			json
// @Produce 		json
// @Param 			id path int64 true "Id"
// @Param 			user body user_dto.UserUpdate true "User Update"
// @Success 		204
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router 			/user/{id} [patch]
func (h *userHandler) Update(c *gin.Context) {
	var in user_dto.UserUpdate

	{
		if err := c.ShouldBindJSON(&in); err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	id, err := statushttp.GetId(c)
	{
		if err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	in.Id = id

	if err = h.srvc.Update(c.Request.Context(), &in); err != nil {
		statushttp.InternalServerError(c, err.Error())
		return
	}

	statushttp.NoContent(c)
}

// GetList 			godoc
// @Summary 		user-service-get-list
// @Description 	user-service-get-list
// @Tags 			User Service
// @Accept 			json
// @Produce 		json
// @Param			page query int true "Page number"
// @Param			limit query int true "Page size"
// @Param 			input query user_dto.UserParams false "Filter parameters"
// @Success 		200 {object} user_dto.UserPage "User list retrieved successfully"
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router 			/user [get]
func (h *userHandler) GetList(c *gin.Context) {
	var in user_dto.UserParams

	{
		if err := c.ShouldBindQuery(&in); err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	page, limit, err := statushttp.GetPageLimit(c)
	{
		if err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	items, err := h.srvc.Page(c.Request.Context(), &in, page, limit)
	{
		if err != nil {
			statushttp.InternalServerError(c, err.Error())
			return
		}
	}

	statushttp.OK(c, items)
}

// Search 			godoc
// @Summary 		user-service-search
// @Description 	user-service-search
// @Tags 			User Service
// @Accept 			json
// @Produce 		json
// @Param 			input query user_dto.UserParams true "Filter parameters"
// @Success 		200 {array}  user_dto.User "User list retrieved successfully"
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router 			/user/search [get]
func (h *userHandler) Search(c *gin.Context) {
	var in user_dto.UserParams

	{
		if err := c.ShouldBindQuery(&in); err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	items, err := h.srvc.Find(c.Request.Context(), &in)
	{
		if err != nil {
			statushttp.InternalServerError(c, err.Error())
			return
		}
	}

	statushttp.OK(c, items)
}

// GetById 			godoc
// @Summary 		user-service-get-by-id
// @Description 	user-service-get-by-id
// @Tags 			User Service
// @Accept 			json
// @Produce 		json
// @Param 			id path int64 true "Id"
// @Success 		200 {object} user_dto.User "User retrieved successfully"
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router 			/user/{id} [get]
func (h *userHandler) GetById(c *gin.Context) {
	id, err := statushttp.GetId(c)
	{
		if err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	item, err := h.srvc.FindOne(c.Request.Context(), &pg.Id{Id: id})
	{
		if err != nil {
			statushttp.InternalServerError(c, err.Error())
			return
		}
	}

	statushttp.OK(c, item)
}

// Delete 			godoc
// @Summary 		user-service-delete
// @Description 	user-service-delete
// @Tags 			User Service
// @Accept 			json
// @Produce 		json
// @Param 			id path int64 true "Id"
// @Success 		204
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router 			/user/{id} [delete]
func (h *userHandler) Delete(c *gin.Context) {
	id, err := statushttp.GetId(c)
	{
		if err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	if err := h.srvc.Delete(c.Request.Context(), &pg.Id{Id: id}); err != nil {
		statushttp.InternalServerError(c, err.Error())
		return
	}

	statushttp.NoContent(c)
}
