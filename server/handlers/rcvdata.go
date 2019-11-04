package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/evalphobia/go-timber/timber"
	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

/*RcvData funcao para encapsular o handler e permitir a passagem do banco por parametro*/
func RcvData(cli *timber.Client, cb *gobreaker.CircuitBreaker) gin.HandlerFunc {
	log.Println("Iniciando Req")

	fn := func(c *gin.Context) {

		_, err := cb.Execute(func() (interface{}, error) {
			resp, err := http.Get("http://localhost:8081/retrievedata")

			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			log.Printf(string(body))
			c.JSON(http.StatusCreated, gin.H{"message": "Processado sem erros."})

			return body, nil
		})
		if err != nil {
			log.Printf("Erro")
		}
		log.Printf(cb.State().String())

	}
	return gin.HandlerFunc(fn)

}

/*RcvData2 funcao para encapsular o handler e permitir a passagem do banco por parametro*/
func RcvData2(cli *timber.Client, cb *gobreaker.CircuitBreaker) gin.HandlerFunc {
	log.Println("Iniciando Req")

	fn := func(c *gin.Context) {

		log.Printf("teste exec")
		_, err := http.Get("http://localhost:8081/retrievedata")
		if err != nil {
			log.Printf(err.Error())
		}
		log.Printf("exec ok")

	}
	return gin.HandlerFunc(fn)

}

/*RtrData funcao para encapsular o handler e permitir a passagem do banco por parametro*/
func RtrData(cli *timber.Client, cb *gobreaker.CircuitBreaker) gin.HandlerFunc {
	log.Println("Iniciando Req")

	fn := func(c *gin.Context) {
		log.Printf("teste exec")
		_, err := http.Get("www.google.com")
		if err != nil {
			log.Printf(err.Error())
		}
		log.Printf("exec ok")
	}
	return gin.HandlerFunc(fn)

}
