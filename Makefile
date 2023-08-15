# App Server
.PHONY: dev
dev:
  docker compose -f docker-compose.yml up --build

# MySQL
.PHONY: init-db
init-db:
  docker compose -f docker-compose.dev.yml run --rm api go run mysql/init_db.go

# Lint, Format
.PHONY: lint
lint: tools
  golangci-lint run ./... --timeout=5m

.PHONY: format
format: tools
  golangci-lint run ./... --fix

.PHONY: migrate
migrate:
  docker compose -f docker-compose.yml run --rm app sh -c 'migrate -path mysql/migrations -database "mysql://docker:password@tcp(mysql:3306)/gin-rest?multiStatements=true" up'


