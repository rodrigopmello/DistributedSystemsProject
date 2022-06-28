# DistributedSystemsProject

Distributed Systems Project (INE410130 - Computação Distribuída - UFSC)

This project implements the circuit breaker pattern in Go. It also implements a remote procedure call, which is employed as an urban environment experiment.

## Usage

The circuit breaker configuration can be defined in the file called cb.json, which establishes the threshold and time to change to the half-open state. RPC is executed in the following host and port: localhost:8080/exec

One can configure the server responsible for the simulation on the file called config.json present in the directory called remotecall. This file is responsible for the server error parameters.


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

There are also commands to build a docker image of this project. 

```
$ make docker-build
$ make docker-run
```





