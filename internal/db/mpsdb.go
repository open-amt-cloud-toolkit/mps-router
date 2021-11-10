/*********************************************************************
 * Copyright (c) Intel Corporation 2021
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package mpsdb

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Device struct {
	GUID        string
	MPSinstance sql.NullString
}

func getDBConnectionStr() string {
	connectionString, ok := os.LookupEnv("MPS_CONNECTION_STRING")
	if !ok {
		log.Println("MPS_CONNECTION_STRING env is not set")
		return ""
	}
	return connectionString
}

func connectToDB(dbSource string) (*sql.DB, error) {
	if dbSource == "" {
		return nil, errors.New("empty db connection string")
	}
	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("failed to open a db connection: ", err)
	}
	return db, nil
}

func getMPSInstance(db *sql.DB, guid string) (string, error) {
	var device Device
	deviceSql := "SELECT guid, mpsinstance FROM devices WHERE guid = $1;"
	if db != nil {
		row := db.QueryRow(deviceSql, guid)
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

func Health() bool {
	dbSource := getDBConnectionStr()
	db, err := connectToDB(dbSource)
	if err != nil {
		log.Println("Failed to open a DB connection: ", err)
		return false
	}
	defer db.Close()
	result := db.QueryRow("SELECT 1")
	if result.Err() != nil {
		log.Println(result.Err().Error())
		return false
	}
	return true
}
func Query(guid string) string {
	dbSource := getDBConnectionStr()
	db, err := connectToDB(dbSource)
	if err != nil {
		log.Println("Failed to open a DB connection: ", err)
		return ""
	}
	defer db.Close()
	mpsInstance := ""
	mpsInstance, err = getMPSInstance(db, guid)
	if err != nil {
		log.Println("Failed to open a DB connection: ", err)
		return ""
	}
	return mpsInstance
}
