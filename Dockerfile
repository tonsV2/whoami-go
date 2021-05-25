FROM golang:1.16-alpine AS build
RUN apk add gcc musl-dev
WORKDIR /src
RUN go get github.com/cespare/reflex
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o /app/whoami-go .

RUN apk --no-cache -U upgrade
FROM alpine:3.9
WORKDIR /app
COPY --from=build /app/whoami-go .
USER guest
CMD ["/app/whoami-go"]
