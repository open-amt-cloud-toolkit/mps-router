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

	p := proxy.NewServer(":8003", "mps:3000")
	log.Println("Proxying from " + p.Addr + " to " + p.Target)
	p.ListenAndServe()
}
