# DistributedSystemsProject

Distributed Systems Project (INE410130 - Computação Distribuída - UFSC)

This project implements the circuit breaker pattern in Go. It is also implemented a remote procedure call, which is employed as part of urban environment experiment. 

## Usage

The circuit breaker configuration can be defined in the file called cb.json, which defines the threadhold and time to change to the half-open state. RPC are executed in the following host and port: localhost:8080/exec

One can configure the server responsible for the simulation on the file called config.json present in the directory remotecall. This file is responsible the server errors parameters.


## Run

```
$ go build -o cb
$ ./cb 
```


## Makefile

There is a simple makefile that can be used instead of go commands. 
```
$ make
```

There's also command to build a docker image of this project. 

```
$ make docker-build
$ make docker-run
```





