FROM golang:alpine3.12

ENV CGO_ENABLED=0

WORKDIR /usr/work 

COPY . .

RUN apk update && apk add curl git ca-certificates && update-ca-certificates
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s && cp ./bin/air /bin/air && rm -rf ./bin

RUN go mod vendor
RUN go mod download

CMD ["air", "-c", ".air.toml"]