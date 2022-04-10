FROM golang:alpine as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN apk update && apk add upx ca-certificates openssl && update-ca-certificates
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o /bin/api-binary ./cmd/example/main.go
RUN upx -9 /bin/api-binary

FROM gcr.io/distroless/static:nonroot
WORKDIR /app/
COPY --from=builder /bin/api-binary /bin/api-binary
COPY --from=builder --chown=nonroot /app/config /app/config
ENTRYPOINT ["/bin/api-binary"]