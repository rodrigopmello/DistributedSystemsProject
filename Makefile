NAME=cb
IMAGENAME=project


all: 
	go build -o $(NAME)
	./$(NAME)

docker-build:
	docker build -t $(IMAGENAME) .


docker-run:
	docker run $(IMAGENAME)
