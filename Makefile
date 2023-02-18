run:
	cd deployments && docker-compose up --build -d

test:
	go clean -testcache
	go test -run= ./internal/service/usecase

cover:
	go test -count=1 -coverprofile=coverage.out ./internal/service/usecase
	go tool cover -html=coverage.out
	rm coverage.out

pprof:
	curl http://localhost:8024/debug/pprof/trace\?seconds\=10 > trace.out
	go tool trace trace.out
	rm trace.out

down:
	cd deployments
	docker-compose down
	docker container prune -f
	docker image prune -f
