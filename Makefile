NAME=main
IMAGENAME=projeto


all: 
	go build -o $(NAME)
	./main

docker-build:
	docker build -t $(IMAGENAME) .


docker-run:
	docker run $(IMAGENAME)
