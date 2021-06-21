package main

import (
	"github.com/Yangiboev/golang-neo4j/api"
	"github.com/Yangiboev/golang-neo4j/config"
	"github.com/Yangiboev/golang-neo4j/pkg/logger"
	"github.com/Yangiboev/golang-neo4j/storage"
)

var (
	log  logger.Logger
	strg storage.StorageI
	cfg  config.Config
)

func initDependencies() {
	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "mongo-golang")

	log.Info("main: Neo4j connection config",
		logger.String("Host", cfg.Neo4jHost),
		logger.Int("Port", cfg.Neo4jPort),
		logger.String("Database", cfg.Neo4jDatabase),
	)
	// neo4jDriver
	// neo4jDriver, err :=

	// mongoConn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoString).SetAuth(credential))

	// if err != nil {
	// 	log.Error("error to connect to mongo database", logger.Error(err))
	// }
	// connDB := mongoConn.Database("mongo-golang")

	// log.Info("Connected to MongoDB", logger.Any("database: ", connDB.Name()))
	// strg = storage.NewProductStorage(connDB)
}

func main() {
	initDependencies()
	server := api.New(api.RouterOptions{
		Config:  cfg,
		Log:     log,
		Storage: strg,
	})

	err := server.Run(cfg.Port)

	if err != nil {
		log.Error("Something went wrong", logger.Error(err))
		panic(err)
	}

}
