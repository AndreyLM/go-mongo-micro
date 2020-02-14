package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

const (
	user     string = "root"
	password string = "example"
	host     string = "localhost"
	port     int32  = 27017
	name     string = "sweatdb"
)

// GetDB gets db connection
func GetDB() (db *mongo.Database, err error) {
	if db != nil {
		return
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", user, password, host, port)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		return
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return
	}

	fmt.Println("Connected to MongoDB")
	db = client.Database(name)

	return
}
