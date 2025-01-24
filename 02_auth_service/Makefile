
migrate:
	goose -dir migrations postgres "postgresql://pguser:pgpassword@127.0.0.1:5432/postgres?sslmode=disable" up

create_users:
	goose -dir migrations create users sql 

start:
	go run app/cmd/main.go --config=./app/config/config.yml