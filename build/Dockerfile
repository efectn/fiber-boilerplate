FROM golang:alpine

ENV CGO_ENABLED=0

WORKDIR /app 

COPY . .

RUN apk update && apk add curl git ca-certificates && update-ca-certificates

RUN go mod download

RUN go build -ldflags "-s -w" -o /app/main cmd/example/main.go
RUN ls /app
CMD ["ls /app"]
#CMD ["/app/main"]
