FROM golang:1.12-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /app/go-sample-app .

FROM alpine:3.9
WORKDIR /app
COPY --from=build /app/go-sample-app .
CMD ["/app/go-sample-app"]
