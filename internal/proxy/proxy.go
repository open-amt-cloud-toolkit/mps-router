/*********************************************************************
 * Copyright (c) Intel Corporation 2021
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package proxy

import (
	"log"
	mpsdb "mps-lookup/internal/db"
	"net"
	"regexp"
	"strings"
)

// Server is a TCP server that takes an incoming request and sends it to another
// server, proxying the response back to the client.
type Server struct {
	// TCP address to listen on
	Addr string
	// TCP address of target server
	Target string

	serve func(ln net.Listener) error
}

func NewServer(addr string, target string) Server {
	if addr == "" {
		addr = ":8003"
	}
	server := Server{
		Addr:   addr,
		Target: target,
	}
	server.serve = server.serveDefault
	return server
}

// ListenAndServe listens on the TCP network address laddr and then handle packets
// on incoming connections.
func (s Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	return s.serve(listener)
}
func (s Server) serveDefault(ln net.Listener) error {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go s.handleConn(conn)
	}
}
func (s Server) parseGuid(content string) string {
	guid := ""
	log.Println("content :", content)
	splitString := strings.Split(content, "\n")
	if len(splitString) < 1 {
		return guid
	}
	r := regexp.MustCompile("[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}")
	guid = r.FindString(splitString[0])
	return guid
}
func (s Server) handleConn(conn net.Conn) {
	destChannel := make(chan net.Conn)
	go s.forward(conn, destChannel)
	dst := <-destChannel
	go s.backward(conn, dst)
}
func (s Server) forward(conn net.Conn, destChannel chan net.Conn) {
	var dst net.Conn
	defer conn.Close()

	buff := make([]byte, 65535)
	isFirst := true
	for {
		n, err := conn.Read(buff)
		if err != nil {
			log.Println(err)
			return
		}
		b := buff[:n]
		destination := s.Target
		if isFirst {
			guid := s.parseGuid(string(b))
			if guid != "" {
				//call to database to get the mps instance
				instance := mpsdb.Query(guid)
				if (instance != "") {
					parts := strings.Split(destination, ":")
					parts[0] = instance
					destination = parts[0] + ":" + parts[1]
				}
			}
			// connects to target server
			dst, err = net.Dial("tcp", destination)
			if err != nil {
				log.Println(err.Error())
				return
			}
			destChannel <- dst
			isFirst = false
		}
		defer dst.Close()
		if err != nil {
			return
		}
		_, err = dst.Write(b)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (s Server) backward(conn net.Conn, dst net.Conn) {
	defer func() {
		conn.Close()
		dst.Close()
	}()
	buff := make([]byte, 65535)
	for {
		n, err := dst.Read(buff)
		if err != nil {
			log.Println(err)
			return
		}
		b := buff[:n]
		_, err = conn.Write(b)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
