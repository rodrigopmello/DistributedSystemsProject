package cb

import (
	"errors"
	"log"
	"time"
)

/*State represents an state(open, closed, and half-open) from the cb*/
type State string

/*Options defines the cb params*/
type Options struct {
	Failurethreshold int
	Retrytimeperiod  time.Duration
}

/*Circuitbreaker represents cb*/
type Circuitbreaker struct {
	o               Options
	s               State
	lastfailuretime time.Time
	failurecount    int
}

var (
	//ErrCircuitBreakerOpen will throw if circuit breaker is open
	ErrCircuitBreakerOpen = errors.New("CB: Circuit Breaker is open")
	//ErrRequest will throw if circuit breaker is closed, but the request failed for some othe reason
	ErrRequest = errors.New("CB: Error during request execution")
	//ErrCircuitBreakerHalfOpen will throw if circuit breaker is half open
	ErrCircuitBreakerHalfOpen = errors.New("CB: Circuit Breaker is half-open")
)

func New(opt Options) *Circuitbreaker {
	var cb Circuitbreaker
	cb.o = opt
	cb.s = "closed"
	return &cb
}

func (c *Circuitbreaker) setState() {
	log.Printf("CB: State definition")
	log.Printf("CB: Failure count %d", c.failurecount)
	if c.failurecount > c.o.Failurethreshold {
		if time.Now().Sub(c.lastfailuretime) > c.o.Retrytimeperiod {
			c.s = "half-open"
		} else {
			c.s = "open"
		}
	} else {
		c.s = "closed"
	}

}

/*CallFunc is responsible to execute the function f*/
func (c *Circuitbreaker) CallFunc(f func() (interface{}, error)) (interface{}, error) {
	c.setState()
	switch c.s {
	case "closed":
		log.Printf("CB: Closed state")
		output, err := f()
		if err != nil {
			c.recordfailure()
			return nil, err
		}
		c.reset()
		return output, nil

	case "open":
		log.Printf("CB: Open state")
		//TODO: return an error
		//c.recordfailure()
		//c.failurecount++
		return nil, ErrCircuitBreakerOpen

	case "half-open":
		log.Printf("CB: Half-open state")
		output, err := f()
		if err != nil {

			c.recordfailure()
			c.s = "open"
			return nil, ErrCircuitBreakerHalfOpen
		}
		c.s = "open"
		c.reset()

		return output, nil

		//return nil, ErrCircuitBreakerHalfOpen

	}

	return nil, nil

}

func (c *Circuitbreaker) setstate() {
	c.reset()

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
