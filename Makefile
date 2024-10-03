build:
	docker build -f compilers/dockerfile.python -t compiler-container-python:latest .
	docker build -f compilers/dockerfile.golang -t compiler-container-golang:latest .


create:
	docker create --name python -v $(CURDIR)/code:/app -t compiler-container-python
	docker create --name golang -v $(CURDIR)/code:/app -t compiler-container-golang

start:
	docker start python
	docker start golang

stop:
	docker stop python
	docker stop golang

delete:
	docker rm golang
	docker rm python

test:
	docker exec -it python python main.py
	docker exec -it golang go run main.go