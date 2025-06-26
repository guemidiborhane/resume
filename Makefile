.PHONY=run
run: main.go
	go run main.go

.PHONY=build
build:
	docker build -t guemidiborhane/resume:latest .
