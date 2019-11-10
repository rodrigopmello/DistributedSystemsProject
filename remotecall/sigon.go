package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

/*Argsagent representa os argumentos necessario para a execucao da funcao de awareness do sigon*/
type Argsagent struct {
	PositionX float64
	PositionY float64
	Speed     float64
	RoadID    int
	Direction string
}

/*SigonAwareness é o agente que irá definir se um usuário de smartphone está consciente ou nao*/
type SigonAwareness struct {
	Musicon             bool
	Carsound            string
	Touchingscreen      bool
	Nearestcardirection string
	Threadhold          float64
	cars                []car
}

type car struct {
	carID     string
	roadID    string
	positionX float64
	positionY float64
}

func (s *SigonAwareness) init() {
	var cars [1]car
	cars[0] = car{"1", "1", 2.0, 2.0}
	fmt.Print("teste")
	fmt.Printf(cars[0].carID)
	s.Musicon = true
	s.Carsound = "high"
	s.Touchingscreen = false

}

/*New instancia do sigon*/
func New() *SigonAwareness {
	var sigon SigonAwareness
	//sigon.init()
	return &sigon
}

func (s *SigonAwareness) samediretion() bool {
	return true
}

/*Notify funcao que ira consultar o sigon para definir nivel de awareness e notificar o pedestre*/
func (s *SigonAwareness) Notify(args *Argsagent, reply *string) error {
	log.Printf("exec")
	if s.distance(args.PositionX, args.PositionY) < s.Threadhold && s.samediretion() {
		*reply = "NotificyCar"
	}
	*reply = "NotifyPedestrian"
	return nil

}

func (s *SigonAwareness) distance(px float64, py float64) float64 {
	/*x := math.Pow((px - s.cars[0].positionX), 2)
	y := math.Pow((py - s.cars[0].positionY), 2)*/

	return 2.0
}

func main() {
	s := New()
	s.init()

	rep := new(SigonAwareness)

	err := rpc.Register(rep)
	if err != nil {
		log.Fatal("Format of service isn't correct. ", err)
	}
	// Register a HTTP handler
	rpc.HandleHTTP()
	// Listen to TPC connections on port 1234
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	log.Printf("Serving RPC server on port %d", 1234)
	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}
