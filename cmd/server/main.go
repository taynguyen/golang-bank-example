package main

import (
	"gin-boilerplate/cmd/server/routers"
	"gin-boilerplate/config"
	"gin-boilerplate/infra/logger"
	"gin-boilerplate/internal/repository"
	"time"

	"github.com/spf13/viper"
)

func main() {

	//set timezone
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Dhaka")
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
	if err := repo.Migrate(); err != nil {
		logger.Fatalf("migrations Migrate() error: %s", err)
	}

	router := routers.SetupRoute(repo)
	logger.Fatalf("%v", router.Run(config.ServerConfig()))
}
