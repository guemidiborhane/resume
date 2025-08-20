PORT?=3000

.PHONY=run
run:
	python -m http.server $(PORT)

.PHONY=build
build:
	docker build -t ghcr.io/guemidiborhane/resume:latest .
