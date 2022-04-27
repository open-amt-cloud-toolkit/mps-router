/*********************************************************************
 * Copyright (c) Intel Corporation 2021
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package main

import (
	"flag"
	"log"
	mpsdb "mps-lookup/internal/db"
	"mps-lookup/internal/proxy"
	"os"
)

func main() {

	result := flag.Bool("health", false, "check health of service")
	flag.Parse()
	if *result {
		dbHealth := mpsdb.Health()
		if dbHealth {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	routerPort := os.Getenv("PORT")
	if routerPort == "" {
		log.Println("PORT env is not set, default is 8003")
		routerPort = "8003"
	}
	mpsPort := os.Getenv("MPS_PORT")
	if mpsPort == "" {
		log.Println("MPS_PORT env is not set, default is 3000")
		mpsPort = "3000"
	}
	mpsHost := os.Getenv("MPS_HOST")
	if mpsHost == "" {
		log.Println("MPS_HOST env is not set,default is mps")
		mpsHost = "mps"
	}

	p := proxy.NewServer(":"+routerPort, mpsHost+":"+mpsPort)
	log.Println("Proxying from " + p.Addr + " to :" + p.Target)
	p.ListenAndServe()
}
