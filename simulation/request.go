package main

import (
	"ProjetoFinalDistribuida/sigon"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

	var param sigon.Argsagent
	param.PositionX = 1
	param.PositionY = 2
	param.Speed = 1
	param.RoadID = 1
	param.Direction = "north"

	//Monta Json
	JSON, err := json.MarshalIndent(param, "", "	")
	if err != nil {
		log.Printf(err.Error())
	}

	for index := 0; index < 10; index++ {
		client := &http.Client{}
		req, err := http.NewRequest("POST", "http://localhost:8080/exec", bytes.NewBuffer(JSON))
		if err != nil {
			log.Printf(err.Error())
		}
		req.Header.Add("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			log.Printf(err.Error())
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		log.Println(string(body))
		time.Sleep(18 * time.Second)

	}

}
