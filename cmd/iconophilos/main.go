package main

import (
	"log"

	"github.com/iconophilos/backend/internal/app"
	mnmtctrl "github.com/iconophilos/backend/internal/pkg/monuments/controller"
	mnmtrepo "github.com/iconophilos/backend/internal/pkg/monuments/repository"
	mnmtsvc "github.com/iconophilos/backend/internal/pkg/monuments/service"
	"github.com/iconophilos/backend/pkg/db"
	envars "github.com/netflix/go-env"
	"go.uber.org/zap"
)

type config struct {
	Env          string `env:"ENV,default=dev"`
	AppName      string `env:"APP_NAME,default=iconophilos"`
	Port         string `env:"PORT,default=8080"`
	AllowHeaders string `env:"ALLOW_HEADERS,default=Origin, Content-Type, Accept"`
	Frontend     string `env:"FRONTEND,default=http://localhost:3000"`
}

func newConfig() *config {
	var cfg config
	if _, err := envars.UnmarshalFromEnviron(&cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg
}

func main() {
	cfg := newConfig()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("could not initialize logger: %s", err)
	}

	defer func() {
		_ = logger.Sync()
	}()

	logger.Named(cfg.AppName)

	logger.Info("getting db connection")

	var dbCfg db.Config
	var dbCfgErr error
	if cfg.Env == "dev" {
		if dbCfg, dbCfgErr = db.NewLocalCfg(); dbCfgErr != nil {
			logger.Fatal("could not initialize local database configuration", zap.Error(dbCfgErr))
		}
	} else {
		if dbCfg, dbCfgErr = db.NewCloudCfg(); dbCfgErr != nil {
			logger.Fatal("could not initialize cloud database configuration", zap.Error(dbCfgErr))
		}
	}

	dbConn, err := db.Connection(dbCfg)
	if err != nil {
		logger.Fatal("failed to get database connection", zap.Error(err))
	}

	// Database migrations

	logger.Info("running db migrations")

	if err := dbConn.AutoMigrate(
		&mnmtrepo.Monument{},
	); err != nil {
		logger.Fatal("failed to apply user database migration", zap.Error(err))
	}

	monumentsRepo := mnmtrepo.NewPostgresRepository(logger, dbConn)
	monumentsSvc := mnmtsvc.NewDefaultService(logger, monumentsRepo)
	monumentsCtrl := mnmtctrl.NewRESTCtrl(logger, monumentsSvc)

	app := app.New(logger, cfg.Port, monumentsCtrl)

	if err := app.Run(); err != nil {
		logger.Fatal("could not start app", zap.Error(err))
	}
}
