package config

import (
	"os"

	"github.com/spf13/cast"
)

const (
	CollectionName = "mac"
)

type Config struct {
	Environment   string
	Neo4jHost     string
	Neo4jPort     int
	Neo4jDatabase string
	Neo4jPassword string
	Neo4jUser     string
	LogLevel      string
	Port          string
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue

}
func Load() Config {

	cfg := Config{}

	cfg.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "develop"))
	cfg.Neo4jHost = cast.ToString(getOrReturnDefaultValue("NEO4J_HOST", "localhost"))
	cfg.Neo4jPort = cast.ToInt(getOrReturnDefaultValue("NEO4J_PORT", 7687))
	cfg.Neo4jDatabase = cast.ToString(getOrReturnDefaultValue("NEO4J_DATABASE", "neo4j"))
	cfg.Neo4jUser = cast.ToString(getOrReturnDefaultValue("NEO4J_USER", "neo4j"))
	cfg.Neo4jPassword = cast.ToString(getOrReturnDefaultValue("NEO4J_PASSWORD", "admin"))
	cfg.LogLevel = cast.ToString(getOrReturnDefaultValue("LOG_LEVEL", "debug"))
	cfg.Port = cast.ToString(getOrReturnDefaultValue("PORT", ":8000"))

	return cfg

}
