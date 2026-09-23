.PHONY: help run run-all run-all-parallel run-all-sequential \
	run-simpleHttpServer run-fmwk-api run-fmwk-test \
	run-lb-api run-lb-lb \
	run-grpc-control run-grpc-worker \
	run-channels run-pubsub run-mutex \
	list

# Auto-discovered main.go files:
#  - concurrencyStuff/channels/main.go
#  - concurrencyStuff/channels/pubsub/main.go
#  - concurrencyStuff/mutex_wg/main.go
#  - fmwkHttpServer/cmd/api/main.go
#  - fmwkHttpServer/cmd/test/main.go
#  - grpc-pg/control-plane/main.go
#  - grpc-pg/worker-agent/main.go
#  - httpLoadBalancer/api/main.go
#  - httpLoadBalancer/loadBalancer/main.go
#  - simpleHttpServer/cmd/api/main.go

help: ## Show this help
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z0-9_.-]+:.*##/ {printf "  %-22s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

list: ## List all discovered main.go files
	@find . -name "main.go" -type f | sort

# ── Individual runners ──────────────────────────────────────────────
run-proxy: ## Run reverseProxy/cmd/api
	@echo "==> reverseProxy/cmd/api"
	@cd reverseProxy && go run ./cmd/api
run-simpleHttpServer: ## Run simpleHttpServer/cmd/api
	@echo "==> simpleHttpServer/cmd/api"
	@cd simpleHttpServer && go run ./cmd/api

run-fmwk-api: ## Run fmwkHttpServer/cmd/api (Fiber :8080, needs DB)
	@echo "==> fmwkHttpServer/cmd/api"
	@cd fmwkHttpServer && go run ./cmd/api

run-fmwk-test: ## Run fmwkHttpServer/cmd/test
	@echo "==> fmwkHttpServer/cmd/test"
	@cd fmwkHttpServer && go run ./cmd/test

run-lb-api: ## Run httpLoadBalancer/api (:3000)
	@echo "==> httpLoadBalancer/api"
	@cd httpLoadBalancer && go run ./api

run-lb-lb: ## Run httpLoadBalancer/loadBalancer (empty main.go)
	@echo "==> httpLoadBalancer/loadBalancer"
	@cd httpLoadBalancer && go run ./loadBalancer

run-grpc-control: ## Run grpc-pg/control-plane (:50051)
	@echo "==> grpc-pg/control-plane"
	@cd grpc-pg && go run ./control-plane

run-grpc-worker: ## Run grpc-pg/worker-agent (client, needs control-plane)
	@echo "==> grpc-pg/worker-agent"
	@cd grpc-pg && go run ./worker-agent

run-channels: ## Run concurrencyStuff/channels
	@echo "==> concurrencyStuff/channels"
	@cd concurrencyStuff/channels && go run main.go

run-pubsub: ## Run concurrencyStuff/channels/pubsub
	@echo "==> concurrencyStuff/channels/pubsub"
	@cd concurrencyStuff/channels/pubsub && go run main.go

run-mutex: ## Run concurrencyStuff/mutex_wg
	@echo "==> concurrencyStuff/mutex_wg"
	@cd concurrencyStuff/mutex_wg && go run main.go

# ── Aggregate runners ───────────────────────────────────────────────

# Sequential: runs one after another. Will block on first long-running
# server (Fiber/gRPC/http). Useful for short-lived demos.
run-all-sequential: run-channels run-pubsub run-mutex ## Run all short-lived demos sequentially
	@echo "Done (servers not included - they block). Use run-all-parallel for servers."

# Run all short-lived demos sequentially (no servers)
run: run-channels run-pubsub run-mutex ## Alias for run-all-sequential (safe, non-blocking)

# Parallel: runs every main.go in background. Servers block, demos exit quickly.
# Ctrl+C kills all background jobs via trap.
run-all-parallel: ## Run ALL main.go files in parallel (background)
	@echo "Running ALL main.go files in parallel (Ctrl+C to stop servers)..."
	@trap 'kill 0; exit' INT TERM; \
	echo "==> concurrencyStuff/channels"; (cd concurrencyStuff/channels && go run main.go) & \
	echo "==> concurrencyStuff/channels/pubsub"; (cd concurrencyStuff/channels/pubsub && go run main.go) & \
	echo "==> concurrencyStuff/mutex_wg"; (cd concurrencyStuff/mutex_wg && go run main.go) & \
	echo "==> fmwkHttpServer/cmd/test"; (cd fmwkHttpServer && go run ./cmd/test) & \
	echo "==> fmwkHttpServer/cmd/api (:8080)"; (cd fmwkHttpServer && go run ./cmd/api) & \
	echo "==> simpleHttpServer/cmd/api (:8080)"; (cd simpleHttpServer && go run ./cmd/api) & \
	echo "==> httpLoadBalancer/api (:3000)"; (cd httpLoadBalancer && go run ./api) & \
	echo "==> httpLoadBalancer/loadBalancer"; (cd httpLoadBalancer && go run ./loadBalancer) & \
	echo "==> grpc-pg/control-plane (:50051)"; (cd grpc-pg && go run ./control-plane) & \
	echo "==> grpc-pg/worker-agent"; sleep 2; (cd grpc-pg && go run ./worker-agent) & \
	wait

# Default aggregate: prefer parallel because sequential blocks on servers.
# For CI / quick check use `make run` (demos only).
run-all: run-all-parallel ## Default: run all in parallel (alias for run-all-parallel)
