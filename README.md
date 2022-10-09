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

# Develop
```sh
docker-compose up dev
```

# Test
```sh
docker-compose up test
```

# HTTP Request
Perform a http request using HTTPie

```sh
http :8080
```

# Helm
## Package
```sh
make helm-package
```

## Upload to repository
```sh
make publish-helm
```
