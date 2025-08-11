package main

import (
	"github.com/Hot-One/monolith/api"
	"github.com/Hot-One/monolith/config"
	"github.com/Hot-One/monolith/pkg/logger"
	postgresConn "github.com/Hot-One/monolith/pkg/postgres"
	"github.com/Hot-One/monolith/service"
	"github.com/Hot-One/monolith/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		cfg         = config.Load()
		loggerLevel = new(string)

		gormConfig = &postgresConn.GormConfig{
			SkipDefaultTransaction: true,
		}
	)

	*loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.New(*loggerLevel, cfg.ServiceName)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			log.Error("Failed to cleanup logger", logger.Error(err))
			return
		}
	}()

	postgres, err := postgresConn.ConnectPostgres(gormConfig, cfg)
	if err != nil {
		log.Error("Failed to connect to PostgreSQL", logger.Error(err))
		return
	}

	var (
		storages = storage.NewStorage(postgres)
		services = service.NewService(storages, &cfg, log)
		server   = api.SetUpRouter(&api.Router{Cfg: &cfg, Log: log, Srvc: services})
	)

	if err := server.Run(":" + cfg.HTTPPort); err != nil {
		log.Error("Failed to run server", logger.Error(err))
		return
	}
}
