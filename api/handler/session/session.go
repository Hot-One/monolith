package session_handler

import (
	statushttp "github.com/Hot-One/monolith/api/status_http"
	"github.com/Hot-One/monolith/config"
	session_dto "github.com/Hot-One/monolith/dto/session"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/pkg/pg"
	"github.com/Hot-One/monolith/service"
	session_service "github.com/Hot-One/monolith/service/session"
	"github.com/gin-gonic/gin"
)

type sessionHandler struct {
	cfg  *config.Config
	log  logger.Logger
	srvc session_service.SessionServiceInterface
}

func NewSessionHandler(group *gin.RouterGroup, srvc service.ServiceInterface, config *config.Config, logger logger.Logger) {
	handler := &sessionHandler{
		cfg:  config,
		log:  logger,
		srvc: srvc.SessionService(),
	}

	session := group.Group("/session")
	{
		session.POST("", handler.Create)
		session.PATCH("/:id", handler.Update)
		session.GET("", handler.GetList)
		session.GET("/:id", handler.GetById)
		session.GET("/search", handler.Search)
		session.DELETE("/:id", handler.Delete)
	}
}

// Create 			godoc
// @Summary 		session-service-create
// @Description 	session-service-create
// @Tags 			Session Service
// @Accept 			json
// @Produce 		json
// @Param 			input body session_dto.SessionCreate true "Session Create"
// @Success 		201 {object} pg.Id "Session created successfully"
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router 			/session [post]
func (h *sessionHandler) Create(c *gin.Context) {
	var in session_dto.SessionCreate

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
// @Summary 		session-service-update
// @Description 	session-service-update
// @Tags 			Session Service
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Id"
// @Param 			input body session_dto.SessionUpdate true "Session Update"
// @Success 		204
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router 			/session/{id} [patch]
func (h *sessionHandler) Update(c *gin.Context) {
	id, err := statushttp.GetId(c)
	{
		if err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	var in session_dto.SessionUpdate
	{
		if err := c.ShouldBindJSON(&in); err != nil {
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
// @Summary 		session-service-get-list
// @Description 	session-service-get-list
// @Tags 			Session Service
// @Accept 			json
// @Produce 		json
// @Param 			page query int false "Page" default(1)
// @Param 			limit query int false "Limit" default(10)
// @Param 			filter query session_dto.SessionParams false "Filter parameters"
// @Success 		200 {object} session_dto.SessionPage "Session List"
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router 			/session [get]
func (h *sessionHandler) GetList(c *gin.Context) {
	page, limit, err := statushttp.GetPageLimit(c)
	{
		if err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	var queryParams session_dto.SessionParams
	{
		if err := c.ShouldBindQuery(&queryParams); err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	items, err := h.srvc.Page(c.Request.Context(), &queryParams, page, limit)
	{
		if err != nil {
			statushttp.InternalServerError(c, err.Error())
			return
		}
	}

	statushttp.OK(c, items)
}

// Search 			godoc
// @Summary 		session-service-search
// @Description 	session-service-search
// @Tags 			Session Service
// @Accept 			json
// @Produce 		json
// @Param 			filter query session_dto.SessionParams false "Filter parameters"
// @Success 		200 {array} session_dto.Session "Session list retrieved successfully"
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router 			/session/search [get]
func (h *sessionHandler) Search(c *gin.Context) {
	var queryParams session_dto.SessionParams
	{
		if err := c.ShouldBindQuery(&queryParams); err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	items, err := h.srvc.Find(c.Request.Context(), &queryParams)
	{
		if err != nil {
			statushttp.InternalServerError(c, err.Error())
			return
		}
	}

	statushttp.OK(c, items)
}

// GetById 			godoc
// @Summary 		session-service-get-by-id
// @Description 	session-service-get-by-id
// @Tags 			Session Service
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Id"
// @Success 		200 {object} session_dto.Session "Session retrieved successfully"
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router 			/session/{id} [get]
func (h *sessionHandler) GetById(c *gin.Context) {
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
// @Summary 		session-service-delete
// @Description 	session-service-delete
// @Tags 			Session Service
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Id"
// @Success 		204
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router 			/session/{id} [delete]
func (h *sessionHandler) Delete(c *gin.Context) {
	id, err := statushttp.GetId(c)
	{
		if err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	if err = h.srvc.Delete(c.Request.Context(), &pg.Id{Id: id}); err != nil {
		statushttp.InternalServerError(c, err.Error())
		return
	}

	statushttp.NoContent(c)
}
