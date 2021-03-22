package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/Yangiboev/golang-postgres-monolith/api"
	"github.com/Yangiboev/golang-postgres-monolith/config"
	"github.com/Yangiboev/golang-postgres-monolith/pkg/logger"
	"github.com/Yangiboev/golang-postgres-monolith/storage"
)

var (
	log  logger.Logger
	cfg  config.Config
	strg storage.StorageI
)

func initDeps() {
	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "api_gateway")

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDB := sqlx.MustConnect("postgres", psqlString)
	strg = storage.NewStoragePg(connDB)

}

func main() {
	initDeps()

	server := api.New(api.Config{
		Storage: strg,
		Logger:  log,
		Cfg:     cfg,
	})

	server.Run(cfg.HTTPPort)
}
