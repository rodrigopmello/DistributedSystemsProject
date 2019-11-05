package main

import (
	"ProjetoFinalDistribuida/models"
	"ProjetoFinalDistribuida/server"
	"os"

	"github.com/evalphobia/go-timber/timber"
	"github.com/sony/gobreaker"
)

const p = "8080"

var cb *gobreaker.CircuitBreaker

func init() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.MaxRequests = 5
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.2
	}

	cb = gobreaker.NewCircuitBreaker(st)
}

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

	r := server.SetupRouter(cli, cb)
	//r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + p)

}
