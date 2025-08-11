package api

import (
	"github.com/Hot-One/monolith/api/docs"
	"github.com/Hot-One/monolith/api/handler"
	user_handler "github.com/Hot-One/monolith/api/handler/user"
	"github.com/Hot-One/monolith/config"
	"github.com/Hot-One/monolith/pkg/logger"
	"github.com/Hot-One/monolith/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type router struct {
	cfg  *config.Config
	log  logger.Logger
	srvc service.ServiceInterface
}

// @title Monolith API
// @version 1.0
// @description API for Monolith application
// @BasePath /v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func SetUpRouter(option *router) *gin.Engine {
	docs.SwaggerInfo.Title = option.cfg.ServiceName
	docs.SwaggerInfo.Schemes = []string{option.cfg.HTTPScheme}

	var (
		r        = gin.Default()
		handlers = handler.NewHandler(option.srvc, option.cfg, option.log)
	)

	r.Use(gin.Recovery(), gin.Logger(), customCORSMiddleware())

	url := ginSwagger.URL("swagger/doc.json")

	v1 := r.Group("/v1", handlers.Ping)
	{
		v1.GET("/ping")
	}

	user_handler.NewUserHandler(v1, option.srvc, option.cfg, option.log)

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
