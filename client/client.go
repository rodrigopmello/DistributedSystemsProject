package main

import (
	"log"
	"net/rpc"
)

type argsagent struct {
	PositionX float64
	PositionY float64
	Speed     float64
	RoadID    int
	Direction string
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	log.Print(client)
	args := argsagent{2.0, 2.0, 2.0, 1, "north"}

	var reply = ""
	err = client.Call("SigonAwareness.Notify", args, &reply)
	log.Print(reply)
	if err != nil {
		log.Print(err.Error())
	}

}
