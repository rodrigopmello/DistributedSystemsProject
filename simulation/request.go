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

	var param sigon.Argsagent
	param.PositionX = 1
	param.PositionY = 2
	param.Speed = 1
	param.RoadID = 1
	param.Direction = "north"
	param.Simulation = true

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

		log.Printf("Req %d Resp: %s", index+1, string(body))
		time.Sleep(18 * time.Second)

	}

}
