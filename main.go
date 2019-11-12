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
	/*err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}*/
	conf := timber.Config{
		APIKey:         os.Getenv("TIMBER_API_KEY"),
		SourceID:       os.Getenv("TIMBER_SOURCE_ID"),
		CustomEndpoint: "https://logs.timber.io",
		Environment:    "production",
		MinimumLevel:   timber.LogLevelFatal,
		Sync:           true,
		Debug:          true,
	}

	cli, err := timber.New(conf)
	models.HandleError(err)
	var opt cb.Options
	opt.Failurethreshold = 1
	dur, err := time.ParseDuration("10s")
	if err != nil {
		log.Println("Error during parsing duration")
		cli.Fatal("S1: Error during duration settup")
	}
	opt.Retrytimeperiod = dur

	//log.Printf("%+v ", opt)

	cb := cb.New(opt)
	cli.Info("S1: Iniciando webservice REST")
	r := server.SetupRouter(cli, cb)
	//r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + p)

}
