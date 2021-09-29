db-migrate:
	go run github.com/prisma/prisma-client-go migrate

db-generate:
	go run github.com/prisma/prisma-client-go generate

run-api:
	go run cmd/api/main.go

test-unit:
	go test -cover ./...