package main

import (
	"github.com/Hot-One/monolith/api"
	"github.com/Hot-One/monolith/config"
	"github.com/Hot-One/monolith/pkg/logger"
	postgresConn "github.com/Hot-One/monolith/pkg/postgres"
	"github.com/Hot-One/monolith/pkg/static"
	"github.com/Hot-One/monolith/service"
	"github.com/Hot-One/monolith/storage"
	"github.com/gin-gonic/gin"
	gormLog "gorm.io/gorm/logger"
)

func main() {
	var (
		cfg         = config.Load()
		loggerLevel = new(string)

		gormConfig = &postgresConn.GormConfig{
			SkipDefaultTransaction: true,
			Logger:                 gormLog.Default.LogMode(gormLog.Info),
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

	if err := postgresConn.Migrate(postgres); err != nil {
		log.Error("Failed to migrate PostgreSQL database", logger.Error(err))
		return
	}

	var (
		storages = storage.NewStorage(postgres)
		services = service.NewService(storages, &cfg, log)
	)

	serverOption := &api.Router{
		Cfg:    &cfg,
		Log:    log,
		Srvc:   services,
		Routes: []*static.Route{},
	}

	var server = api.SetUpRouter(serverOption)

	if err := postgresConn.CreateSystemRows(storages, serverOption.Routes); err != nil {
		log.Error("Failed to create system rows", logger.Error(err))
		return
	}

	if err := server.Run(":" + cfg.HTTPPort); err != nil {
		log.Error("Failed to run server", logger.Error(err))
		return
	}
}
