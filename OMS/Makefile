# Variables
IMAGE_NAME = oms
CONTAINER_NAME = oms
DOCKERFILE = dockerfile

# Comandos
build:
	sudo docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .

docker-oms:
	sudo docker run -p 8080:80 $(CONTAINER_NAME)
stop:
	sudo docker stop $(CONTAINER_NAME)
rm:
	sudo docker rm $(CONTAINER_NAME)
clean:	sudo stop rm
	sudo docker rmi $(IMAGE_NAME)
logs:
	sudo docker logs -f $(CONTAINER_NAME)
