build:
	docker build -f compilers/dockerfile.python -t compiler-container-python:latest .
	docker build -f compilers/dockerfile.golang -t compiler-container-golang:latest .


create:
	docker create --name python -v $(CURDIR)/code:/app -t compiler-container-python python main.py
	docker create --name golang -v $(CURDIR)/code:/app -t compiler-container-golang go run main.go

delete:
	docker rm golang
	docker rm python

test:
	docker start -i python
	docker start -i golang