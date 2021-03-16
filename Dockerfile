FROM golang:1.16-alpine AS build
RUN apk add gcc musl-dev
WORKDIR /src
COPY . .
RUN go build -o /app/whoami-go .

FROM alpine:3.9
WORKDIR /app
COPY --from=build /app/whoami-go .
CMD ["/app/whoami-go"]
