package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	Username string
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	log.Print(client)
	args := Args{"rodrigopmello"}

	var reply = ""
	err = client.Call("Findrepositories.Search", args, &reply)
	log.Print(reply)
	if err != nil {
		log.Print(err.Error())
	}

}
