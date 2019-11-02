package server

import (
	"ProjetoFinalDistribuida/server/handlers"

	"github.com/evalphobia/go-timber/timber"
	"github.com/gin-gonic/gin"
)

/*SetupRouter inicializa o gin*/
func SetupRouter(cli *timber.Client) *gin.Engine {
	r := gin.Default()

	r.POST("/rcvdata", handlers.RcvData(cli))

	return r
}
