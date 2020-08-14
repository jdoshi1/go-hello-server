.PHONY: *

build:
	docker build --no-cache -t go-hello-server .

run:
	docker run --rm=true -p 8080:8080 -ti go-hello-server

