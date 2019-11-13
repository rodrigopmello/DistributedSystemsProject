package sigon

import (
	"errors"
	"fmt"
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
		car{
			carID:     "1",
			roadID:    "1",
			positionX: 2.0,
			positionY: 2.0,
		},
	}
	//var cars [1]car
	//cars[0] = car{"1", "1", 2.0, 2.0}
	fmt.Println("teste")
	fmt.Println(s.cars[0].carID)
	s.Musicon = true
	s.Carsound = "high"
	s.Touchingscreen = false
	count2fail1 = count1
	count2fail2 = count2
	log.Printf("%d", count1)
	log.Printf("%d", count2)

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

/*Notify funcao que ira consultar o sigon para definir nivel de awareness e notificar o pedestre*/
func (s *Awareness) Notify(args *Argsagent, reply *string) error {
	count++
	log.Printf("exec %d", count)
	log.Printf("%d ", count2fail1)
	log.Printf("%d ", count2fail2)
	if count == count2fail1 || count == count2fail2 {
		s.previous = time.Now().Add(time.Duration(secondsDown) * time.Second)
	}
	for time.Now().Before(s.previous) {
		return ErrProgramao
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

/*func main() {

	conf := timber.Config{
		APIKey:         os.Getenv("TIMBER_API_KEY"),
		SourceID:       os.Getenv("TIMBER_SOURCE_ID"),
		CustomEndpoint: "https://logs.timber.io",
		Environment:    "production",
		MinimumLevel:   timber.LogLevelError,
		Sync:           true,
		Debug:          true,
	}

	cli, errT := timber.New(conf)

	if errT != nil {
		cli.Info(errT.Error())
	}

	rep := new(Awareness)
	rep.init()

	err := rpc.Register(rep)
	if err != nil {
		log.Fatal("Format of service isn't correct. ", err)
		cli.Fatal("S2: Formato do servico incorreto")
	}
	// Register a HTTP handler
	rpc.HandleHTTP()
	// Listen to TPC connections on port 1234
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error: ", e)
		cli.Fatal("S2: Erro na criacao da conexao TCP" + e.Error())

	}
	log.Printf("Serving RPC server on port %d", 1234)
	cli.Info("S2: Iniciando RPC Server ")
	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
		//cli.Fatal("S2: Erro ao iniciar o RPC Server")
	}

}
*/
