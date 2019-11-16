package main

import (
	"ProjetoFinalDistribuida/cb"
	"ProjetoFinalDistribuida/models"
	"ProjetoFinalDistribuida/server"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/evalphobia/go-timber/timber"
)

const p = "8080"

type cbCfg struct {
	Failurethreshold int    `json:"failurethreshold"`
	Retrytimeperiod  string `json:"retrytimeperiod"`
}

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

	jsonFile, err := os.Open("cb.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var cfg cbCfg

	json.Unmarshal(byteValue, &cfg)

	cli, err := timber.New(conf)
	models.HandleError(err)
	var opt cb.Options
	opt.Failurethreshold = cfg.Failurethreshold
	dur, err := time.ParseDuration(cfg.Retrytimeperiod)
	if err != nil {
		log.Println("Error during parsing duration")
		cli.Fatal("S1: Error during duration setup")
	}
	opt.Retrytimeperiod = dur

	log.Printf("S1: Configurações do Circuit Breaker %+v ", opt)

	cb := cb.New(opt)
	cli.Info("S1: Iniciando webservice REST")
	r := server.SetupRouter(cli, cb)
	//r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + p)

}
