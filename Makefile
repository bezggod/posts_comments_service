migrate-up:
	goose -dir ./db/migrations postgres "host=localhost user=bezgo dbname=p_c_service password=12345 sslmode=disable" up

migrate-create:
	goose -dir ./db/migrations create $(F) sql

run-pg:
		POSTGRES_PG_DSN="postgres://bezgo:12345@localhost:5432/p_c_service?sslmode=disable" \
    	STORAGE_MODE=postgres ADDR=:8080 go run ./cmd