
migrate:
	migrate -path ./migrations -database 'postgres://pguser:pgpassword@localhost:5432/postgres?sslmode=disable' up

start:
	 go run cmd/main.go --config=./config/config.yml