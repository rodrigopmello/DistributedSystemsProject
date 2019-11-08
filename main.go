package main

import (
	"ProjetoFinalDistribuida/cb"
	"ProjetoFinalDistribuida/models"
	"ProjetoFinalDistribuida/server"
	"log"
	"os"
	"time"

	"github.com/evalphobia/go-timber/timber"
)

const p = "8080"

func main() {

	conf := timber.Config{
		APIKey:         os.Getenv("TIMBER_API_KEY"),
		SourceID:       os.Getenv("TIMBER_SOURCE_ID"),
		CustomEndpoint: "https://logs.timber.io",
		Environment:    "production",
		MinimumLevel:   timber.LogLevelInfo,
		Sync:           true,
		Debug:          true,
	}

	cli, err := timber.New(conf)
	models.HandleError(err)
	var opt cb.Options
	opt.Failurethreshold = 3
	opt.Timeout = 3

	opt.Resettimeout = 3
	dur, err := time.ParseDuration("30s")
	if err != nil {
		log.Println("Error during parse duration")
	}
	opt.Retrytimeperiod = dur

	log.Printf("%+v ", opt)

	cb := cb.New(opt)

	r := server.SetupRouter(cli, cb)
	//r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + p)

}
