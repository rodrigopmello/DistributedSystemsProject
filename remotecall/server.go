package main

import (
	"ProjetoFinalDistribuida/sigon"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"

	"github.com/evalphobia/go-timber/timber"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conf := timber.Config{
		APIKey:         os.Getenv("TIMBER_API_KEY"),
		SourceID:       os.Getenv("TIMBER_SOURCE_ID"),
		CustomEndpoint: "https://logs.timber.io",
		Environment:    "production",
		MinimumLevel:   timber.LogLevelInfo,
		Sync:           true,
		Debug:          true,
	}

	cli, errT := timber.New(conf)

	if errT != nil {
		cli.Info(errT.Error())
	}

	rep := new(sigon.Awareness)

	err = rpc.Register(rep)
	if err != nil {
		log.Fatal("Format of service isn't correct. ", err)
		cli.Fatal("S2: Formato do servico incorreto")
	}
	// Register a HTTP handler
	rpc.HandleHTTP()
	// Listen to TPC connections on port 1234
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error: ", e)
		cli.Fatal("S2: Erro na criacao da conexao TCP" + e.Error())

	}
	log.Printf("Serving RPC server on port %d", 1234)
	cli.Info("S2: Iniciando RPC Server ")
	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
		cli.Fatal("S2: Erro ao iniciar o RPC Server")
	}
}
