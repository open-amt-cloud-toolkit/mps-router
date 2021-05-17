/*********************************************************************
 * Copyright (c) Intel Corporation 2021
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package main

import (
	"flag"
	"log"
	"mps-lookup/internal/proxy"
)

func main() {
	flag.Parse()
	p := proxy.Server{
		Addr: ":8003",
	}
	log.Println("Proxying from " + p.Addr + " to :3000" + p.Target)
	p.ListenAndServe()
}
