package cb

type cbOptions struct {
	errorthresholdpercentage int
	timeout                  int
	resettimeout             int
}

/*Circuitbreaker representa o cb*/
type Circuitbreaker struct {
	o cbOptions
}

/*New criacao de um ponteiro para o cb*/
func New(c cbOptions) *Circuitbreaker {
	var cb Circuitbreaker
	cb.o = c
	return &cb
}
