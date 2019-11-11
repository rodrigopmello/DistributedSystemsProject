package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {

	/*{
		 "positionx": 2,
		 "positiony": 2,
	    "speed": 10,
	    "roadid": 1,
	    "direction": "north"

	}
	ticker := time.NewTicker(5 * time.Second)
	for _ = range ticker.C {
		fmt.Println("tock")
		resp, err := http.PostForm("http://localhost:8080/exec",
			url.Values{"positionx": {"2"}, "positiony": {"2"}, "speed": {"2"}, "roadid": {"2"}, "direction": {"north"}})

		if err != nil {
			log.Printf("Erro ao executar requisição " + err.Error())
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		log.Printf(string(body))

	}*/

	for index := 0; index < 10; index++ {
		resp, err := http.PostForm("http://localhost:8080/exec",
			url.Values{"positionx": {"2"}, "positiony": {"2"}, "speed": {"2"}, "roadid": {"2"}, "direction": {"north"}})

		if err != nil {
			log.Printf("Erro ao executar requisição " + err.Error())
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		log.Printf(string(body))
		time.Sleep(5 * time.Second)

	}

}
