package db

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoManager This code is responsible for communicating with the MongoDB compatible database and
// is used to retrieve and store data to and from the database.
type MongoManager struct {
	// ConnectionString is the connection string to the MongoDB compatible database.
	ConnectionString string
	// DatabaseName is the name of the database to use. Default is "mpsdb"
	DatabaseName string
	// CollectionName is the name of the collection to use. Default is "devices"
	CollectionName string
}

func NewMongoManager(connectionString string) *MongoManager {
	databaseName := os.Getenv("MPS_DATABASE_NAME")
	if databaseName == "" {
		databaseName = "mpsdb"
	}
	collectionName := os.Getenv("MPS_COLLECTION_NAME")
	if collectionName == "" {
		collectionName = "devices"
	}
	return &MongoManager{
		ConnectionString: connectionString,
		DatabaseName:     databaseName,
		CollectionName:   collectionName,
	}
}

func (m *MongoManager) Connect() (Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.ConnectionString))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (m *MongoManager) GetMPSInstance(db Database, guid string) (string, error) {
	return "", errors.New("not implemented")
}

func (m *MongoManager) Health() bool {
	client, err := m.Connect()
	if err != nil {
		log.Println(err.Error())
		return false
	}

	// We'll cast our generic Database type back to a *mongo.Client.
	mongoClient, ok := client.(*mongo.Client)
	if !ok {
		return false
	}
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Printf("Failed to disconnect from db: %v", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Using Ping to check the health.
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func (m *MongoManager) Query(guid string) string {
	client, err := m.Connect()
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	// We'll cast our generic Database type back to a *mongo.Client.
	mongoClient, ok := client.(*mongo.Client)
	if !ok {
		return ""
	}
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Printf("Failed to disconnect from db: %v", err)
		}
	}()

	// Using the same logic as in GetMPSInstance to fetch the MPSinstance.
	collection := mongoClient.Database(m.DatabaseName).Collection(m.CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var device Device
	err = collection.FindOne(ctx, map[string]interface{}{"guid": guid}).Decode(&device)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return device.MPSinstance
}
