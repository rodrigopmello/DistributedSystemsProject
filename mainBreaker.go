package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker

func init() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.2
	}

	cb = gobreaker.NewCircuitBreaker(st)
}

// Get wraps http.Get in CircuitBreaker.
func Get(url string) ([]byte, error) {
	body, err := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err

		}

		return body, nil
	})

	log.Printf(cb.State().String())
	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}

func main() {
	for index := 0; index < 1000; index++ {
		log.Printf("executando")
		time.Sleep(5 * time.Second)
		body, err := Get("http://localhost:8080/rcvdata")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))
	}

}
