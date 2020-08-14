go-hello-server
=====

Simple go server with readiness and health endpoints using gorilla/mux for routing.

# Pre-reqs

1. Make sure to install and run docker daemon
    - [docker-desktop](https://docs.docker.com/desktop/) for Mac/Windows or
    - [docker-engine](https://docs.docker.com/engine/) for Linux

# How to build locally in docker

```bash
make build
```

# How to run the server using docker

```bash
make run
```

# How to test

On a different terminal window run:

```bash
curl localhost:8080
curl localhost:8080?name=jdoshi1
curl localhost:8080/health
curl localhost:8080/readiness
curl localhost:8080/blah
```

# How to shutdown the server

Press ctrl+c