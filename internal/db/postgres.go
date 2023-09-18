/*********************************************************************
 * Copyright (c) Intel Corporation 2021
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package db

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

type PostgresManager struct {
	ConnectionString string
}

func (pm *PostgresManager) Connect() (Database, error) {
	db, err := sql.Open("postgres", pm.ConnectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (pm *PostgresManager) GetMPSInstance(db Database, guid string) (string, error) {
	client, ok := db.(*sql.DB)
	if !ok {
		return "", errors.New("invalid database type for PostgreSQL")
	}
	var device Device
	deviceSql := "SELECT guid, mpsinstance FROM devices WHERE guid = $1;"
	if client != nil {
		row := client.QueryRow(deviceSql, guid)
		switch err := row.Scan(&device.GUID, &device.MPSinstance); err {
		case sql.ErrNoRows:
			log.Println("no rows were returned!")
		case nil:
			{
				return device.MPSinstance.String, nil
			}
		default:
			{
				log.Println("failed to execute query: ", err)
				return "", err
			}
		}
		return "", nil
	}
	return "", errors.New("invalid db connection")
}

func (pm *PostgresManager) Health() bool {
	db, err := pm.Connect()
	if err != nil {
		log.Println("Failed to open a DB connection: ", err)
		return false
	}

	defer db.(*sql.DB).Close()
	result := db.(*sql.DB).QueryRow("SELECT 1")
	if result.Err() != nil {
		log.Println(result.Err().Error())
		return false
	}
	return true
}
func (pm *PostgresManager) Query(guid string) string {
	db, err := pm.Connect()
	if err != nil {
		log.Println("Failed to open a DB connection: ", err)
		return ""
	}
	defer db.(*sql.DB).Close()
	mpsInstance := ""
	mpsInstance, err = pm.GetMPSInstance(db, guid)
	if err != nil {
		log.Println("Failed to open a DB connection: ", err)
		return ""
	}
	return mpsInstance
}
