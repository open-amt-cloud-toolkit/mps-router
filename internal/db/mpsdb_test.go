/*********************************************************************
 * Copyright (c) Intel Corporation 2021
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package mpsdb

import (
	"database/sql"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBConfigWithNoEnvironmentVariables(t *testing.T) {
	result := getDBConnectionStr()
	assert.Equal(t, "", result)
}
func TestDBConfigWithEnvironmentVariables(t *testing.T) {
	// Set an Environment Variable
	os.Setenv("MPS_CONNECTION_STRING", "postgresql://postgresadmin:admin123@localhost:5432/mpsdb?sslmode=disable")

	expected := "postgresql://postgresadmin:admin123@localhost:5432/mpsdb?sslmode=disable"
	result := getDBConnectionStr()
	assert.Equal(t, expected, result)
}
func TestConnectToDBWithValidConnectionString(t *testing.T) {
	var db *sql.DB
	result, err := connectToDB("postgresql://")
	assert.Nil(t, err, "test failed to connect db")
	assert.Equal(t, reflect.TypeOf(result), reflect.TypeOf(db))
}
func TestConnectToDBWithInvalidConnectionString(t *testing.T) {
	_, err := connectToDB("")
	assert.Equal(t, "empty db connection string", err.Error())
}

func TestGetMPSInstancewithGUID(t *testing.T) {
	db, err := connectToDB("postgresql://")
	assert.Nil(t, err, "test failed to connect db")
	result := ""
	result, err = getMPSInstance(db, "d12428be-9fa1-4226-9784-54b2038beab6")
	if err != nil {
		log.Println("test failed to get the mps instance", err)
	}
	assert.Equal(t, "", result)
}

func TestGetMPSInstancewithNoDB(t *testing.T) {
	var db *sql.DB
	_, err := getMPSInstance(db, "d12428be-9fa1-4226-9784-54b2038beab6")
	if err != nil {
		log.Println("test failed to get the mps instance", err)
	}
	assert.Equal(t, "invalid db connection", err.Error())
}

func TestQuery(t *testing.T) {
	// Set an Environment Variable
	os.Setenv("MPS_CONNECTION_STRING", "postgresql://")
	result := Query("d12428be-9fa1-4226-9784-54b2038beab6")
	assert.Equal(t, "", result)
}
