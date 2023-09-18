package test

import (
	"database/sql"
	"mps-lookup/internal/db"
)

type MockDBManager struct {
	ConnectResult     *sql.DB
	ConnectError      error
	ConnectionStr     string
	MPSInstanceResult string
	MPSInstanceError  error
	HealthResult      bool
	QueryResult       string
}

func (mock *MockDBManager) Connect() (db.Database, error) {
	return mock.ConnectResult, mock.ConnectError
}
func (mock *MockDBManager) GetMPSInstance(db db.Database, guid string) (string, error) {
	if mock.MPSInstanceError != nil {
		return "", mock.MPSInstanceError
	}
	return mock.MPSInstanceResult, nil
}

func (mock *MockDBManager) Health() bool {
	return mock.HealthResult
}

func (mock *MockDBManager) Query(guid string) string {
	return mock.QueryResult
}
