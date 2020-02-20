package db

import (
	"context"
	"fmt"
	"time"

	"github.com/andreylm/go-mongo-micro/sqmplemgr/logger"

	"github.com/andreylm/go-mongo-micro/sqmplemgr/config"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

// Init - inits db connection
func Init() {
	host := config.ReadEnvString("DB_HOST")
	port := config.ReadEnvInt("DB_PORT")
	name := config.ReadEnvString("DB_NAME")
	user := config.ReadEnvString("DB_USER")
	password := config.ReadEnvString("DB_PASSWORD")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", user, password, host, port)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		logger.Get().Fatal("Cannot initialize database")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		logger.Get().Fatal("Cannot initialize database context")
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Get().Fatal("Cannot ping database")
	}

	logger.Get().Info("Connected To MongoDB")
	db = client.Database(name)
	return
}

// GetDB gets db connection
func GetDB() *mongo.Database {
	if db == nil {
		logger.Get().Fatal("Database not initialized")
		return nil
	}

	return db
}
