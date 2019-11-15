package server

import (
	"ProjetoFinalDistribuida/cb"
	"ProjetoFinalDistribuida/server/handlers"

	"github.com/evalphobia/go-timber/timber"
	"github.com/gin-gonic/gin"
)

/*SetupRouter inicializa o gin*/
func SetupRouter(cli *timber.Client, cb *cb.Circuitbreaker) *gin.Engine {

	r := gin.Default()

	r.GET("/rcvdata", handlers.RcvData(cli, cb))
	r.GET("/rcvdata2", handlers.RcvData2(cli, cb))
	r.POST("/exec", handlers.ExecRemoteCall(cli, cb))
	r.POST("/execSimulation", handlers.ExecRemoteCallSimulation(cli, cb))

	r.GET("/retrievedata", handlers.RtrData(cli, cb))

	return r
}
