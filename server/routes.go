package server

import (
	"ProjetoFinalDistribuida/server/handlers"

	"github.com/evalphobia/go-timber/timber"
	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

/*SetupRouter inicializa o gin*/
func SetupRouter(cli *timber.Client, cb *gobreaker.CircuitBreaker) *gin.Engine {
	r := gin.Default()

	r.GET("/rcvdata", handlers.RcvData(cli, cb))
	r.GET("/rcvdata2", handlers.RcvData2(cli, cb))

	r.GET("/retrievedata", handlers.RtrData(cli, cb))

	return r
}
