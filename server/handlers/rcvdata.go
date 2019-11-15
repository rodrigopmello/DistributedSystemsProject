package handlers

import (
	"ProjetoFinalDistribuida/cb"
	"ProjetoFinalDistribuida/sigon"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"

	"github.com/evalphobia/go-timber/timber"
	"github.com/gin-gonic/gin"
)

/*RcvData funcao para encapsular o handler e permitir a passagem do banco por parametro*/
func RcvData(cli *timber.Client, cb *cb.Circuitbreaker) gin.HandlerFunc {
	log.Println("Iniciando Req")

	fn := func(c *gin.Context) {

		_, err := cb.CallFunc(func() (interface{}, error) {
			resp, err := http.Get("http://localhost:8081/retrievedata")

			if err != nil {
				log.Printf("teste auxiliar %s", err.Error())
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
		//log.Printf(cb.State().String())

	}
	return gin.HandlerFunc(fn)

}

/*RcvData2 funcao para encapsular o handler e permitir a passagem do banco por parametro*/
func RcvData2(cli *timber.Client, cb *cb.Circuitbreaker) gin.HandlerFunc {
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
func RtrData(cli *timber.Client, cb *cb.Circuitbreaker) gin.HandlerFunc {
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

type Args struct {
	Username string
}

/*ExecRemoteCall funcao que ira executar uma remote call usando o circuit breaker*/
func ExecRemoteCall(cli *timber.Client, cb *cb.Circuitbreaker) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		log.Println("Iniciando Remote Call")
		cli.Info("S1: Executando remote procudure call (Awareness.Notify)")
		_, err := cb.CallFunc(func() (interface{}, error) {

			client, err := rpc.DialHTTP("tcp", "localhost:1234")
			if err != nil {
				log.Print("Connection error: ", err)
				cli.Fatal("S1: Erro de conexão")
				return nil, err
			}
			defer client.Close()
			log.Print(client)
			var param sigon.Argsagent
			err1 := c.BindJSON(&param)
			if err1 != nil {
				log.Printf(err1.Error())
				cli.Err(err1.Error())
				cli.Err("S1: Erro de binding do json")

				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return nil, err1
			}

			var reply = ""
			err = client.Call("Awareness.Notify", param, &reply)

			if err != nil {
				log.Print(err.Error())
				cli.Err(err.Error())
				c.JSON(http.StatusCreated, gin.H{"message": "NotifyTrafficLight"})
				return nil, err
			}
			log.Print(reply)
			cli.Info("S1: Respostas de S2: " + reply)
			c.JSON(http.StatusCreated, gin.H{"message": reply})
			return nil, nil
		})
		if err != nil {
			log.Printf("Erro")
			cli.Err(err.Error())
		}

		//log.Printf(cb.State().String())

	}
	return gin.HandlerFunc(fn)

}
