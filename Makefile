PORT:=8080
REGISTRY:=github.com/lucasvmiguel
API_IMAGE:=go-api-test
VERSION:=latest

db-migrate:
	go run github.com/prisma/prisma-client-go migrate

db-generate:
	go run github.com/prisma/prisma-client-go generate

run-api:
	go run cmd/api/main.go

build-api: db-migrate db-generate
	go build cmd/api/main.go

test-unit:
	go test -cover ./...

docker-build-api:
	docker build -t $(REGISTRY)/$(API_IMAGE):$(VERSION) -f cmd/api/Dockerfile .

docker-run-api:
	docker run --rm -p $(PORT):$(PORT) $(REGISTRY)/$(API_IMAGE):$(VERSION)
