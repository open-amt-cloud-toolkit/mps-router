package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestConnect(t *testing.T) {

	manager := NewMongoManager("mongodb://localhost:27017")
	db, err := manager.Connect()
	assert.Nil(t, err, "test failed to connect db")

	_, ok := db.(*mongo.Client)
	assert.Equal(t, true, ok)
}

func TestNoSQLHealth(t *testing.T) {
	manager := &MongoManager{
		ConnectionString: "mongodb://localhost:27017",
	}
	result := manager.Health()
	assert.Equal(t, false, result)
}

func TestNoSQLGetMPSInstance(t *testing.T) {
	manager := &MongoManager{
		ConnectionString: "mongodb://localhost:27017",
	}
	result, err := manager.GetMPSInstance(nil, "mockGUID")
	assert.Empty(t, result)
	assert.Error(t, err)
}
func TestNoSQLQuery(t *testing.T) {
	manager := &MongoManager{
		ConnectionString: "mongodb://localhost:27017",
	}
	res := manager.Query("mockGUID")
	assert.Empty(t, res)
}
