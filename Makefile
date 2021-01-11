setup: teardown run_services
	@go run ./cmd/db/setup.go

run_services:
	@docker-compose up --build -d

run_server:
	@KUDOS_PORT=4444 KUDOS_HOST=localhost go run cmd/main.go

run_client:
	@/bin/bash -c "cd ./pkg/http/web/app && yarn serve"

teardown:
	@docker-compose down --remove-orphans