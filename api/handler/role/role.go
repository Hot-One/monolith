package role_handler

import (
	statushttp "github.com/Hot-One/monolith/api/status_http"
	"github.com/Hot-One/monolith/config"
	role_dto "github.com/Hot-One/monolith/dto/role"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/pkg/pg"
	"github.com/Hot-One/monolith/service"
	role_service "github.com/Hot-One/monolith/service/role"
	"github.com/gin-gonic/gin"
)

type roleHandler struct {
	cfg  *config.Config
	log  logger.Logger
	srvc role_service.RoleServiceInterface
}

func NewRoleHandler(group *gin.RouterGroup, srvc service.ServiceInterface, cfg *config.Config, log logger.Logger) {
	handler := &roleHandler{
		cfg:  cfg,
		log:  log,
		srvc: srvc.RoleService(),
	}

	role := group.Group("/role")
	role.Use(srvc.AuthService().Middleware())
	{
		role.POST("", handler.Create)
		role.PATCH("/:id", handler.Update)
		role.GET("", handler.GetList)
		role.GET("/:id", handler.GetById)
		role.GET("/search", handler.Search)
		role.DELETE("/:id", handler.Delete)
	}
}

// Create 			godoc
// @Security		BearerAuth
// @ID 				role-service-create
// @Summary     	role-service-create
// @Description 	role-service-create
// @Tags        	Role Service
// @Accept      	json
// @Produce     	json
// @Param 			input body role_dto.RoleCreate true "Role Create"
// @Success 		201 {object} pg.Id "Role created successfully"
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router      	/role [post]
func (h *roleHandler) Create(c *gin.Context) {
	var in role_dto.RoleCreate

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

// Update 			godoc
// @Security		BearerAuth
// @ID 				role-service-update
// @Summary			role-service-update
// @Description		role-service-update
// @Tags			Role Service
// @Accept			json
// @Produce	    	json
// @Param 			id path int64 true "Id"
// @Param			input body role_dto.RoleUpdate true "Role Update"
// @Success			204
// @Failure			400 {object} statushttp.Response "Bad Request"
// @Failure			500 {object} statushttp.Response "Internal Server Error"
// @Router			/role/{id} [patch]
func (h *roleHandler) Update(c *gin.Context) {
	id, err := statushttp.GetId(c)
	{
		if err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	var in role_dto.RoleUpdate
	{
		if err := c.ShouldBindJSON(&in); err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	in.Id = id

	if err := h.srvc.Update(c.Request.Context(), &in); err != nil {
		statushttp.InternalServerError(c, err.Error())
		return
	}

	statushttp.NoContent(c)
}

// GetList 			godoc
// @Security		BearerAuth
// @ID 				role-service-getlist
// @Summary 		role-service-getlist
// @Description 	role-service-getlist
// @Tags        	Role Service
// @Accept     		json
// @Produce    		json
// @Param			page query int true "Page number"
// @Param			limit query int true "Page size"
// @Param 			input query role_dto.RoleParams false "Filter parameters"
// @Success    		200 {object} role_dto.Role "Role List"
// @Failure    		400 {object} statushttp.Response "Bad Request"
// @Failure    		500 {object} statushttp.Response "Internal Server Error"
// @Router     		/role [get]
func (h *roleHandler) GetList(c *gin.Context) {
	var in role_dto.RoleParams

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
// @Security		BearerAuth
// @ID 				role-service-search
// @Summary			role-service-search
// @Description		role-service-search
// @Tags			Role Service
// @Accept			json
// @Produce			json
// @Param 			input query role_dto.RoleParams true "Filter parameters"
// @Success 		200 {array}  role_dto.Role "Role list retrieved successfully"
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router 			/role/search [get]
func (h *roleHandler) Search(c *gin.Context) {
	var in role_dto.RoleParams

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

// GetById			godoc
// @Security		BearerAuth
// @ID 				role-service-getbyid
// @Summary			role-service-getbyid
// @Description		role-service-getbyid
// @Tags			Role Service
// @Accept			json
// @Produce			json
// @Param 			id path int64 true "Id"
// @Success 		200 {object} role_dto.Role "Role retrieved successfully"
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router 			/role/{id} [get]
func (h *roleHandler) GetById(c *gin.Context) {
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
// @Security		BearerAuth
// @ID 				role-service-delete
// @Summary			role-service-delete
// @Description		role-service-delete
// @Tags			Role Service
// @Accept			json
// @Produce			json
// @Param			id path int64 true "Id"
// @Success	204	
// @Failure	400 	{object} statushttp.Response "Bad Request"
// @Failure	500 	{object} statushttp.Response "Internal Server Error"
// @Router			/role/{id} [delete]
func (h *roleHandler) Delete(c *gin.Context) {
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
