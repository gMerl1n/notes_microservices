
migrate:
	migrate -path ./app/migrations -database 'postgres://pguser:pgpassword@localhost:5432/postgres?sslmode=disable' up

start:
	go run app/cmd/main.go --config=./app/config/config.yml