package app

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	mnmtctrl "github.com/iconophilos/backend/internal/pkg/monuments/controller"
)

type App struct {
	logger   *zap.Logger
	port     string
	router   *gin.Engine
	mnmtCtrl mnmtctrl.Controller
}

func New(logger *zap.Logger, port string, mnmtCtrl mnmtctrl.Controller) *App {
	return &App{
		logger:   logger,
		port:     port,
		router:   gin.Default(),
		mnmtCtrl: mnmtCtrl,
	}
}

func (a *App) Run() error {
	a.logger.Info("starting app...", zap.String("port", a.port))

	// Middleware

	a.router.Use(gin.Logger())
	a.router.Use(gin.Recovery())

	// Routes

	v1 := a.router.Group("/v1")
	{
		// Monuments

		v1.POST("/monuments", a.mnmtCtrl.Create)
		v1.DELETE("/monuments/:id", a.mnmtCtrl.Delete)
		v1.GET("/monuments", a.mnmtCtrl.List)
		v1.GET("/monuments/:id", a.mnmtCtrl.FetchByID)
	}

	if err := a.router.Run(":" + a.port); err != nil {
		return fmt.Errorf("could not run http server: %s", err)
	}
	return nil
}
