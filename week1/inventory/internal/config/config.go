package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI    string
	MongoDBName string
	GRPCPort    string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load(".env")

	user := os.Getenv("MONGO_USER")
	pass := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	dbName := os.Getenv("MONGO_DB_NAME")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, pass, host, port)

	return &Config{
		MongoURI: uri,
		MongoDBName: dbName,
		GRPCPort:    os.Getenv("GRPC_PORT"),
	}, nil

}
