package db

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoManager struct {
	ConnectionString string
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
	client, ok := db.(*mongo.Client)
	if !ok {
		return "", errors.New("invalid database type for MongoDB")
	}

	collection := client.Database("mpsdb").Collection("devices")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var device struct {
		GUID        string `bson:"guid"`
		MPSinstance string `bson:"mpsinstance"`
	}

	err := collection.FindOne(ctx, map[string]interface{}{"guid": guid}).Decode(&device)
	if err != nil {
		return "", err
	}

	return device.MPSinstance, nil
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
	defer mongoClient.Disconnect(context.Background())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Using Ping to check the health.
	err = mongoClient.Ping(ctx, nil)

	return err != nil
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
	defer mongoClient.Disconnect(context.Background())

	// Using the same logic as in GetMPSInstance to fetch the MPSinstance.
	collection := mongoClient.Database("mpsdb").Collection("devices")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var device struct {
		GUID        string `bson:"guid"`
		MPSinstance string `bson:"mpsinstance"`
	}

	err = collection.FindOne(ctx, map[string]interface{}{"guid": guid}).Decode(&device)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No documents match the filter.
			return ""
		}
		return ""
	}

	return device.MPSinstance
}
