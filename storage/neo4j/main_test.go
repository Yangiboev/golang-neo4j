package neo4j_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Yangiboev/golang-neo4j/config"
	"github.com/Yangiboev/golang-neo4j/pkg/logger"
	"github.com/Yangiboev/golang-neo4j/storage"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var (
	log  logger.Logger
	cfg  config.Config
	strg storage.StorageI
)

func TestMain(m *testing.M) {
	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "golang-neo4j")
	log.Info("main: Neo4j connection config",
		logger.String("Host", cfg.Neo4jHost),
		logger.Int("Port", cfg.Neo4jPort),
		logger.String("Database", cfg.Neo4jDatabase),
	)
	connUrl := fmt.Sprintf("neo4j://%s:%d", cfg.Neo4jHost, cfg.Neo4jPort)
	neoDriver, err := neo4j.NewDriver(connUrl, neo4j.BasicAuth(cfg.Neo4jUser, cfg.Neo4jPassword, ""))
	defer neoDriver.Close()
	if err != nil {
		log.Error("error to connect to neo4j", logger.Error(err))
		panic(err)
	}
	log.Info("Connected to Neo4j", logger.Any("neo4j: ", cfg.Neo4jDatabase))
	strg = storage.NewResponsibleStorage(neoDriver)
	os.Exit(m.Run())
}
