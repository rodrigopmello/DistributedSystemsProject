package sigon

import (
	"fmt"
	"log"
)

/*Argsagent representa os argumentos necessario para a execucao da funcao de awareness do sigon*/
type Argsagent struct {
	PositionX float64
	PositionY float64
	Speed     float64
	RoadID    int
	Direction string
}

/*Awareness é o agente que irá definir se um usuário de smartphone está consciente ou nao*/
type Awareness struct {
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

func (s *Awareness) init() {
	var cars [1]car
	cars[0] = car{"1", "1", 2.0, 2.0}
	fmt.Print("teste")
	fmt.Printf(cars[0].carID)
	s.Musicon = true
	s.Carsound = "high"
	s.Touchingscreen = false

}

/*New instancia do sigon*/
func New() *Awareness {
	var sigon Awareness
	//sigon.init()
	return &sigon
}

func (s *Awareness) samediretion() bool {
	return true
}

/*Notify funcao que ira consultar o sigon para definir nivel de awareness e notificar o pedestre*/
func (s *Awareness) Notify(args *Argsagent, reply *string) error {
	log.Printf("exec")
	if s.distance(args.PositionX, args.PositionY) < s.Threadhold && s.samediretion() {
		*reply = "NotificyCar"
	}
	*reply = "NotifyPedestrian"
	return nil

}

func (s *Awareness) distance(px float64, py float64) float64 {
	/*x := math.Pow((px - s.cars[0].positionX), 2)
	y := math.Pow((py - s.cars[0].positionY), 2)*/

	return 2.0
}
