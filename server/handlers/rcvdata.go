package handlers

import (
	"log"
	"net/http"

	"github.com/evalphobia/go-timber/timber"
	"github.com/gin-gonic/gin"
)

/*RcvData funcao para encapsular o handler e permitir a passagem do banco por parametro*/
func RcvData(cli *timber.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		log.Println("Iniciando IH_010")
		cli.Info("Iniciando IH_010")

		c.JSON(http.StatusCreated, gin.H{"message": "Processado sem erros."})
	}
	return gin.HandlerFunc(fn)

}
