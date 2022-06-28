package sigon

import (
	"errors"
	"log"
	"math"
	"time"
)

var count = 0

var (
	//ErrProgramao will throw if circuit breaker is half open
	ErrProgramao = errors.New("Erro programado")
)

/*Argsagent representa os argumentos necessario para a execucao da funcao de awareness do sigon*/
type Argsagent struct {
	PositionX  float64
	PositionY  float64
	Speed      float64
	RoadID     int
	Direction  string
	Simulation bool //parametro para controlar se é uma simulacao ou nao
}

/*Awareness é o agente que irá definir se um usuário de smartphone está consciente ou nao*/
type Awareness struct {
	Musicon             bool
	Carsound            string
	Touchingscreen      bool
	Nearestcardirection string
	Threadhold          float64
	cars                []car
	count               int
	previous            time.Time
}

/*Simulation variables*/

var count2fail1 int
var count2fail2 int
var secondsDown int

type car struct {
	carID     string
	roadID    string
	positionX float64
	positionY float64
}

//Init funcao para definir parametros dummies para o sigon
func (s *Awareness) Init(count1 int, count2 int, secondsOff int) {
	s.cars = []car{
		{
			carID:     "1",
			roadID:    "1",
			positionX: 2.0,
			positionY: 2.0,
		},
	}
	//var cars [1]car
	//cars[0] = car{"1", "1", 2.0, 2.0}
	/*Test params*/
	s.Musicon = true
	s.Carsound = "high"
	s.Touchingscreen = false
	count2fail1 = count1
	count2fail2 = count2
	secondsDown = secondsOff

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

/*Notify function used to defined awareness level and if required, a notification can be triggered.*/
func (s *Awareness) Notify(args *Argsagent, reply *string) error {
	log.Printf("S2: Simulation %t \n", args.Simulation)
	if args.Simulation {
		count++
		log.Printf("S2: Execution %d", count)
		if count == count2fail1 || count == count2fail2 {
			s.previous = time.Now().Add(time.Duration(secondsDown) * time.Second)
		}
		for time.Now().Before(s.previous) {
			return ErrProgramao
		}
	}

	if s.distance(args.PositionX, args.PositionY) < s.Threadhold && s.samediretion() {
		*reply = "NotificyCar"
	}
	*reply = "NotifyPedestrian"

	return nil

}

func (s *Awareness) distance(px float64, py float64) float64 {
	x := math.Pow((px - s.cars[0].positionX), 2)
	y := math.Pow((py - s.cars[0].positionY), 2)

	return math.Sqrt(x + y)
}
