package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

/*Args representa os argumentos necessario para a execucao da funcao*/
type Args struct {
	Username string
}

/*Findrepositories responsavel por executar a funcao*/
type Findrepositories string

/*Search funcao que ira executar a busca na API*/
func (f *Findrepositories) Search(args *Args, reply *string) error {
	log.Printf("exec")
	*reply = "not found"
	return nil
}

func main() {
	rep := new(Findrepositories)

	err := rpc.Register(rep)
	if err != nil {
		log.Fatal("Format of service Task isn't correct. ", err)
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
