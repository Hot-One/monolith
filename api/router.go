package api

import (
	"github.com/Hot-One/monolith/api/docs"
	app_handler "github.com/Hot-One/monolith/api/handler/app"
	auth_handler "github.com/Hot-One/monolith/api/handler/auth"
	role_handler "github.com/Hot-One/monolith/api/handler/role"
	session_handler "github.com/Hot-One/monolith/api/handler/session"
	user_handler "github.com/Hot-One/monolith/api/handler/user"
	"github.com/Hot-One/monolith/config"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/pkg/static"
	"github.com/Hot-One/monolith/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	Cfg    *config.Config
	Log    logger.Logger
	Routes []*static.Route
	Srvc   service.ServiceInterface
}

// @title Monolith API
// @version 1.0
// @description API for Monolith application
// @BasePath /v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func SetUpRouter(option *Router) *gin.Engine {
	docs.SwaggerInfo.Title = option.Cfg.ServiceName
	docs.SwaggerInfo.Schemes = []string{option.Cfg.HTTPScheme}

	var r = gin.Default()
	var url = ginSwagger.URL("/swagger/doc.json")

	r.Use(gin.Recovery(), gin.Logger(), customCORSMiddleware())

	var v1 = r.Group("/v1")

	user_handler.NewUserHandler(v1, option.Srvc, option.Cfg, option.Log)
	role_handler.NewRoleHandler(v1, option.Srvc, option.Cfg, option.Log)
	app_handler.NewAppHandler(v1, option.Srvc, option.Cfg, option.Log)
	session_handler.NewSessionHandler(v1, option.Srvc, option.Cfg, option.Log)
	auth_handler.NewAuthHandler(v1, option.Cfg, option.Log, option.Srvc)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "3600")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
