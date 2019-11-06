package cb

import (
	"log"
	"time"
)

type options struct {
	failurethreshold int
	timeout          int
	resettimeout     int
	retrytimeperiod  time.Duration
}

/*State representa o estado (open, closed, and half-open) do cb*/
type State string

/*Circuitbreaker representa o cb*/
type Circuitbreaker struct {
	o               options
	s               State
	lastfailuretime time.Time
	failurecount    int
}

/*class CircuitBreaker {
	constructor(timeout, failureThreshold, retryTimePeriod) {
	  // We start in a closed state hoping that everything is fine
	  this.state = 'CLOSED';
	  // Number of failures we receive from the depended service before we change the state to 'OPEN'
	  this.failureThreshold = failureThreshold;
	  // Timeout for the API request.
	  this.timeout = timeout;
	  // Time period after which a fresh request be made to the dependent
	  // service to check if service is up.
	  this.retryTimePeriod = retryTimePeriod;
	  this.lastFailureTime = null;
	  this.failureCount = 0;
	}
  }*/

/*New criacao de um ponteiro para o cb*/
func New(opt options) *Circuitbreaker {
	var cb Circuitbreaker
	cb.o = opt
	cb.s = "closed"
	return &cb
}

func (c *Circuitbreaker) setState() {
	if c.failurecount > c.o.failurethreshold {
		if time.Now().Sub(c.lastfailuretime) > c.o.retrytimeperiod {
			c.s = "half-open"
		} else {
			c.s = "open"
		}
	} else {
		c.s = "closed"
	}

}

/*CallFunc funcao responsável por executar a funcao f*/
func (c *Circuitbreaker) CallFunc(f func() (interface{}, error)) (interface{}, error) {
	c.setState()
	switch c.s {
	case "closed":
		output, err := f()
		if err != nil {
			//execucao falhou mesmo assim
			return nil, err
		}
		c.reset()
		log.Printf("Closed state")
		return output, nil

	case "open":
		log.Printf("Open state")
		//TODO: return an error
		return nil, nil

	case "half-open":
		log.Printf("Half-open state")
		//what should i do in this state?

	}

	return nil, nil

}

func (c *Circuitbreaker) reset() {
	c.lastfailuretime = time.Time{}
	c.failurecount = 0
	c.s = "closed"

}

func (c *Circuitbreaker) recordfailure() {
	c.failurecount++
	c.lastfailuretime = time.Now()

}
