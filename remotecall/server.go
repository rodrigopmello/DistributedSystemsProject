package main

import (
	"ProjetoFinalDistribuida/sigon"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	rep := new(sigon.Awareness)

	err := rpc.Register(rep)
	if err != nil {
		log.Fatal("Format of service isn't correct. ", err)
	}
	// Register a HTTP handler
	rpc.HandleHTTP()
	// Listen to TPC connections on port 1234
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	log.Printf("Serving RPC server on port %d", 1234)
	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}
