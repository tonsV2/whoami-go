# Introduction
Simple Go "Who Am I" microservice

Executing an HTTP request to / will return the hostname of the server

# Launch
## Go
```sh
go build -o whoami-go . && ./whoami-go
```

## Docker-compose
```sh
docker-compose up prod
```

# HTTP Request
Perform a http request using HTTPie

```sh
http :8080
```
