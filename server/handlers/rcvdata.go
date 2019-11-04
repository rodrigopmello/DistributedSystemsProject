package handlers

import (
	"log"
	"net/http"

	"github.com/evalphobia/go-timber/timber"
	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker

func init() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	cb = gobreaker.NewCircuitBreaker(st)
}

/*RcvData funcao para encapsular o handler e permitir a passagem do banco por parametro*/
func RcvData(cli *timber.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/*_, err := cb.Execute(func() (interface{}, error) {

		})
		if err != nil {

		}*/
		log.Println("Iniciando IH_010")
		cli.Info("Iniciando IH_010")

		c.JSON(http.StatusCreated, gin.H{"message": "Processado sem erros."})
	}
	return gin.HandlerFunc(fn)

}
