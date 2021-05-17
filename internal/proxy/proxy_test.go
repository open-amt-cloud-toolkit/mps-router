/*********************************************************************
 * Copyright (c) Intel Corporation 2021
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package proxy

import (
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGuidEmptyString(t *testing.T) {
	server := Server{}
	result := server.parseGuid("")
	assert.Equal(t, "", result)
}
func TestParseGuidRequestWithInvalidGUID_HTTP(t *testing.T) {
	server := Server{}
	request := `GET /api/v1/amt/log/audit/12345?startIndex=0 HTTP/1.1 
	User-Agent: PostmanRuntime/7.28.0
	Accept: */*
	Postman-Token: 63f32fee-238e-4f6a-a091-092270d22439
	Host: 127.0.0.1:8003
	Accept-Encoding: gzip, deflate, br
	Connection: keep-alive`
	result := server.parseGuid(request)
	assert.Equal(t, "", result)
}

func TestParseGuidRequestWithNoGUID_HTTP(t *testing.T) {
	server := Server{}
	request := `GET /api/v1/devices HTTP/1.1
	Host: mpslookup:8003
	Connection: keep-alive
	X-Forwarded-For: 10.0.0.2
	X-Forwarded-Proto: https
	X-Forwarded-Host: localhost
	X-Forwarded-Port: 8443
	X-Forwarded-Path: /mps/api/v1/devices
	X-Forwarded-Prefix: /mps
	X-Real-IP: 10.0.0.2
	sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="90", "Google Chrome";v="90"
	accept: application/json, text/plain, */*`
	result := server.parseGuid(request)
	assert.Equal(t, "", result)
}

func TestParseGuidRequestWithGUID_HTTP(t *testing.T) {
	server := Server{}
	request := `GET /api/v1/amt/log/audit/63f32fee-238e-4f6a-a091-092270d22439?startIndex=0 HTTP/1.1 
	User-Agent: PostmanRuntime/7.28.0
	Accept: */*
	Postman-Token: 63f32fee-238e-4f6a-a091-092270d22439
	Host: 127.0.0.1:8003
	Accept-Encoding: gzip, deflate, br
	Connection: keep-alive`
	result := server.parseGuid(request)
	assert.Equal(t, "63f32fee-238e-4f6a-a091-092270d22439", result)
}

func TestParseGuidRequestWithGUID_WS(t *testing.T) {
	server := Server{}
	request := `GET /relay/webrelay.ashx?p=2&host=d12428be-9fa1-4226-9784-54b2038beab6&port=16994&tls=0&tls1only=0 HTTP/1.1
	Host: mpslookup:8003
	Upgrade: websocket
	Connection: keep-alive, Upgrade
	X-Forwarded-For: 10.0.0.2
	X-Forwarded-Proto: https
	X-Forwarded-Host: localhost
	X-Forwarded-Port: 8443
	X-Forwarded-Path: /mps/ws/relay/webrelay.ashx
	X-Forwarded-Prefix: /mps/ws
	X-Real-IP: 10.0.0.2
	Pragma: no-cache
	Cache-Control: no-cache
	User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36`
	result := server.parseGuid(request)
	assert.Equal(t, "d12428be-9fa1-4226-9784-54b2038beab6", result)
}

func TestParseGuidRequestWithInvalidGUID_WS(t *testing.T) {
	server := Server{}
	request := `GET /relay/webrelay.ashx?p=2&host=d12428be-9fa1-4226-9784&port=16994&tls=0&tls1only=0 HTTP/1.1
	Host: mpslookup:8003
	Upgrade: websocket
	Connection: keep-alive, Upgrade
	X-Forwarded-For: 10.0.0.2
	X-Forwarded-Proto: https
	X-Forwarded-Host: localhost
	X-Forwarded-Port: 8443
	X-Forwarded-Path: /mps/ws/relay/webrelay.ashx
	X-Forwarded-Prefix: /mps/ws
	X-Real-IP: 10.0.0.2
	Pragma: no-cache
	Cache-Control: no-cache
	User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36`
	result := server.parseGuid(request)
	assert.Equal(t, "", result)
}

func TestListenAndServe(t *testing.T) {
	server := Server{Addr: ":8009"}
	go func() {
		server.ListenAndServe()
		conn, err := net.Dial("tcp", ":8009")
		if err != nil {
			log.Println(err)
		}
		log.Println("connection :", conn)
		defer conn.Close()
	}()
}
