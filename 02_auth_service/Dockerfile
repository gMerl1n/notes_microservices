FROM golang:1.22-alpine AS builder

WORKDIR /auth_service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -mod=readonly -o auth_service app/cmd/main.go

#./app/cmd/main.go

# Run the binary when the container starts

# CMD ["./auth_service/app/cmd/main"]

# CMD ["/auth_service"]