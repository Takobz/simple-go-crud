module simple-go-crud/notes-api

go 1.23.4

replace simple-go-crud/configuration => ./internal/configuration

require simple-go-crud/repository v0.0.0-00010101000000-000000000000

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.2 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	simple-go-crud/configuration v0.0.0-00010101000000-000000000000 // indirect
)

replace simple-go-crud/repository => ./internal/repository
