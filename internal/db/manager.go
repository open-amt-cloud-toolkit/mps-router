package db

import "database/sql"

// Universal Database object. For MongoDB, we'd use a *mongo.Client, and for SQL, we'd use *sql.DB
type Database interface{}

type Manager interface {
	Connect() (Database, error)
	GetMPSInstance(db Database, guid string) (string, error)
	Health() bool
	Query(guid string) string
}
type Device struct {
	GUID        string
	MPSinstance sql.NullString
}
