FROM golang:1.23.4-alpine AS builder

WORKDIR /auth_service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -mod=readonly -o auth_service main.go

#./app/cmd/main.go

# Run the binary when the container starts

CMD ["./auth_service/main"]

# CMD ["/auth_service"]