package main

import (
	"gin-boilerplate/cmd/server/routers"
	"gin-boilerplate/config"
	"gin-boilerplate/infra/logger"
	accCtrlPkg "gin-boilerplate/internal/controllers/accounts"
	"gin-boilerplate/internal/handlers/accounts"
	"gin-boilerplate/internal/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	//set timezone
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}
	masterDSN, replicaDSN := config.DbConfiguration()

	repo, err := repository.New(masterDSN, replicaDSN)
	if err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	// TODO: later separate migration
	logger.Infof("migrations Migrate() started")
	if err := repo.Migrate(); err != nil {
		logger.Fatalf("migrations Migrate() error: %s", err)
	}
	logger.Infof("migrations Migrate() completed")

	accCtrl := accCtrlPkg.New(repo)
	accHandler := accounts.NewHandler(accCtrl)
	router := routers.NewRouter(accHandler)

	ginEngine := gin.New()
	router.RegisterRoutes(ginEngine)
	logger.Fatalf("%v", ginEngine.Run(config.ServerConfig()))
}
