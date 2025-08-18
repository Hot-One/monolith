package app_handler

import (
	statushttp "github.com/Hot-One/monolith/api/status_http"
	"github.com/Hot-One/monolith/config"
	app_dto "github.com/Hot-One/monolith/dto/app"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/pkg/pg"
	"github.com/Hot-One/monolith/service"
	app_service "github.com/Hot-One/monolith/service/app"
	"github.com/gin-gonic/gin"
)

type appHandler struct {
	cfg  *config.Config
	log  logger.Logger
	srvc app_service.ApplicationServiceInterface
}

func NewAppHandler(group *gin.RouterGroup, srvc service.ServiceInterface, cfg *config.Config, log logger.Logger) {
	handler := &appHandler{
		cfg:  cfg,
		log:  log,
		srvc: srvc.ApplicationService(),
	}

	app := group.Group("/app")
	{
		app.POST("", handler.Create)
	}

}

// Create 			godoc
// @Summary  		app-service-create
// @Description 	app-service-create
// @Tags      		app-service
// @Accept    		json
// @Produce   		json
// @Param     		input  body  app_service.CreateRequest  true  "Create App"
// @Success 		201 {object} pg.Id "Role created successfully"
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router    		/app [post]
func (h *appHandler) Create(c *gin.Context) {
	var in app_dto.ApplicationCreate
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
// @Summary  		app-service-update
// @Description 	app-service-update
// @Tags      		app-service
// @Accept    		json
// @Produce   		json
// @Param     		id     path  string  true  "Id"
// @Param     		input  body  app_service.UpdateRequest  true  "Update App"
// @Success 		204
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router    		/app [put]
func (h *appHandler) Update(c *gin.Context) {
	id, err := statushttp.GetId(c)
	{
		if err != nil {
			statushttp.BadRequest(c, err.Error())
			return
		}
	}

	var in app_dto.ApplicationUpdate
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
// @Summary  		app-service-get-list
// @Description 	app-service-get-list
// @Tags      		app-service
// @Accept    		json
// @Produce   		json
// @Param     		page  query  int  false  "Page number" default(1)
// @Param     		limit query  int  false  "Number of items per page" default(10)
// @Param 			input query user_dto.UserParams false "Filter parameters"
// @Success 		200 {object} app_dto.ApplicationPage "List of Apps"
// @Failure 		400 {object} statushttp.Response "Bad Request"
// @Failure 		500 {object} statushttp.Response "Internal Server Error"
// @Router    		/app [get]
func (h *appHandler) GetList(c *gin.Context) {
	var in app_dto.ApplicationParams
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

	items, err := h.srvc.Page(c.Request.Context(), page, limit, &in)
	{
		if err != nil {
			statushttp.InternalServerError(c, err.Error())
			return
		}
	}

	statushttp.OK(c, items)
}

// Search 			godoc
// @Summary 		app-service-search
// @Description 	app-service-search
// @Tags        	app-service
// @Accept     		json
// @Produce    		json
// @Param			input query app_dto.ApplicationSearchParams false "Search parameters"
// @Success    		200 {array}  app_dto.Application "Search Results"
// @Failure    		400 {object} statushttp.Response "Bad Request"
// @Failure    		500 {object} statushttp.Response "Internal Server Error"
// @Router     		/app/search [get]
func (h *appHandler) Search(c *gin.Context) {
	var in app_dto.ApplicationParams
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
// @Summary 		app-service-get-by-id
// @Description 	app-service-get-by-id
// @Tags        	app-service
// @Accept     		json
// @Produce    		json
// @Param			id path string true "Id"
// @Success    		200 {object} app_dto.Application "Application retrieved successfully"
// @Failure    		400 {object} statushttp.Response "Bad Request"
// @Failure    		500 {object} statushttp.Response "Internal Server Error"
// @Router     		/app/{id} [get]
func (h *appHandler) GetById(c *gin.Context) {
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
// @Summary 		app-service-delete
// @Description 	app-service-delete
// @Tags        	app-service
// @Accept     		json
// @Produce    		json
// @Param			id path string true "Id"
// @Success    		204
// @Failure    		400 {object} statushttp.Response "Bad Request"
// @Failure    		500 {object} statushttp.Response "Internal Server Error"
// @Router     		/app/{id} [delete]
func (h *appHandler) Delete(c *gin.Context) {
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
