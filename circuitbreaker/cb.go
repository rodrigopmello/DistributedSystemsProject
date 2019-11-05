package cb

type options struct {
	errorthresholdpercentage int
	timeout                  int
	resettimeout             int
}

type State int

/*Circuitbreaker representa o cb*/
type Circuitbreaker struct {
	o options
	s State
}

/*New criacao de um ponteiro para o cb*/
func New(c options) *Circuitbreaker {
	var cb Circuitbreaker
	cb.o = c
	return &cb
}

func (c *Circuitbreaker) currentState() {

}

func (c *Circuitbreaker) changeState() {

}

func (c *Circuitbreaker) resetTimer() {

}
