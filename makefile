# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# CLASS NOTES
#
# Kind
# 	For full Kind v0.26 release notes: https://github.com/kubernetes-sigs/kind/releases/tag/v0.26.0
#
# RSA Keys
# 	To generate a private/public key PEM file.
# 	$ openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
# 	$ openssl rsa -pubout -in private.pem -out public.pem
# 	$ ./admin genkey
#

run:
	go run ./api/services/sales/main.go | go run ./api/tooling/logfmt/main.go

help:
	go run ./api/services/sales/main.go --help

version:
	go run ./api/services/sales/main.go --version
	
curl-live:
	curl -il -X GET http://localhost:3000/v1/liveness

curl-ready:
	curl -il -X GET http://localhost:3000/v1/readiness

curl-test-error:
	curl -il -X GET http://localhost:3000/v1/test-error

curl-test-panic:
	curl -il -X GET http://localhost:3000/v1/test-panic

genKeys:
	go run ./api/tooling/admin/main.go

# admin token (paste on terminal to register as an environment variable)
# export TOKEN=eyJhbGciOiJSUzI1NiIsImtpZCI6IjU0YmIyMTY1LTcxZTEtNDFhNi1hZjNlLTdkYTRhMGUxZTJjMSIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzZXJ2aWNlIHByb2plY3QiLCJzdWIiOiJiZGVhZjE3Yy05MDc0LTRhZmUtOWJlMS1iNDU3YjJiOTU0MTQiLCJleHAiOjE3NzM1NTI5NDcsImlhdCI6MTc0MjAxNjk0Nywicm9sZXMiOlsiQURNSU4iXX0.ZMfzXzmxSVEm3xR25hvL2VMjW9U_23L0mU6PnpPLK8ST88MPIfWEmRnvmyNJ62JcbgjYKq3s1uMQGd58v0gLvw-KkvcY6UMdE9XJ-O07_cAo8E1hWB5zkNA4uWj8UU3KNAE2zMPzEwHxpar_PMRwJHRpxot41yXk8_qG6M1nfrMTyY2P-rQP-0aySbH_3Zh6TJ7zSGde35lGPDpdAFTo9m-SUoXbmSeV2iFMZShK8fTfRnsG5ciPIFqYHWDpmSflPSZEkpU6pecvYPWM20TaFrDFJVPQ2WOGYdO7T3DM53A1irS0BS3iZul_CBPGurvpBsdMaUANDts1563HL54aNQ

# user token
# export TOKEN=eyJhbGciOiJSUzI1NiIsImtpZCI6IjU0YmIyMTY1LTcxZTEtNDFhNi1hZjNlLTdkYTRhMGUxZTJjMSIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzZXJ2aWNlIHByb2plY3QiLCJzdWIiOiJiZGVhZjE3Yy05MDc0LTRhZmUtOWJlMS1iNDU3YjJiOTU0MTQiLCJleHAiOjE3NzM1NTMwNDAsImlhdCI6MTc0MjAxNzA0MCwicm9sZXMiOlsiVVNFUiJdfQ.rj05gkPsPVe7ejQz0lnFCm1Yy-8MFETw3vCcvsIxcMFL6skZgizc_DC3Ii9Qy2pyUj4TN-GvGItmBqrwJ7h4lQD-NwqOo9dZAOzmm6qNxXvLqzBT_0HeTwPD1flnCciZN0U-5J8i9iEMxreE6Wo0wTajk46hNBVx3JyoZTORzqasIZPRcAzMvFweqdftGSBnE-GyAA-72Ggh42Gzi6-GXmUpOTpRPzfECZBhWvp0gq5dquRTu09Am5ulGxj02fm30Au6qU2HoBft0oLStDdrqHze7ujpBPz_jeGeLooOpXHYamVzNjaA5PpyjzJmDs-KNLkvCq7Hy1IJRkb84cLJoQ

curl-test-auth:
	curl -i \
	-H "Authorization: Bearer ${TOKEN}" "http://localhost:3000/v1/test-auth"

curl-token:
	curl -i \
	--user "admin@example.com:gophers" http://localhost:6000/v1/auth/token/54bb2165-71e1-41a6-af3e-7da4a0e1e2c1

curl-authen:
	curl -i \
	-H "Authorization: Bearer ${TOKEN}" "http://localhost:6000/v1/auth/authenticate"

# ==============================================================================
# Define dependencies

GOLANG          := golang:1.24
ALPINE          := alpine:3.21
KIND            := kindest/node:v1.32.0
POSTGRES        := postgres:17.3
GRAFANA         := grafana/grafana:11.5.0
PROMETHEUS      := prom/prometheus:v3.1.0
TEMPO           := grafana/tempo:2.7.0
LOKI            := grafana/loki:3.4.0
PROMTAIL        := grafana/promtail:3.4.0

KIND_CLUSTER    := fbdaf-starter-cluster
NAMESPACE       := sales-system
SALES_APP       := sales
AUTH_APP        := auth
BASE_IMAGE_NAME := localhost/fbdaf
VERSION         := 0.0.1
SALES_IMAGE     := $(BASE_IMAGE_NAME)/$(SALES_APP):$(VERSION)
METRICS_IMAGE   := $(BASE_IMAGE_NAME)/metrics:$(VERSION)
AUTH_IMAGE      := $(BASE_IMAGE_NAME)/$(AUTH_APP):$(VERSION)

# ==============================================================================
# Building containers

build: sales auth

sales:
	docker build \
		-f zarf/docker/dockerfile.sales \
		-t $(SALES_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.

auth:
	docker build \
		-f zarf/docker/dockerfile.auth \
		-t $(AUTH_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.

# ==============================================================================
# Running from within k8s/kind

dev-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/dev/kind-config.yaml

	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

	kind load docker-image $(POSTGRES) --name $(KIND_CLUSTER) 

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)

dev-status-all:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

dev-status:
	watch -n 2 kubectl get pods -o wide --all-namespaces

# ------------------------------------------------------------------------------

dev-load:
	kind load docker-image $(SALES_IMAGE) --name $(KIND_CLUSTER) & \
	kind load docker-image $(AUTH_IMAGE) --name $(KIND_CLUSTER) & \
	wait;

dev-load-db:
	kind load docker-image $(POSTGRES) --name $(KIND_CLUSTER)

dev-apply:
	kustomize build zarf/k8s/dev/database | kubectl apply -f -
	kubectl rollout status --namespace=$(NAMESPACE) --watch --timeout=120s sts/database

	kustomize build zarf/k8s/dev/auth | kubectl apply -f -
	kubectl wait pods --namespace=$(NAMESPACE) --selector app=$(AUTH_APP) --timeout=120s --for=condition=Ready

	kustomize build zarf/k8s/dev/sales | kubectl apply -f -
	kubectl wait pods --namespace=$(NAMESPACE) --selector app=$(SALES_APP) --timeout=120s --for=condition=Ready

dev-restart:
	kubectl rollout restart deployment $(AUTH_APP) --namespace=$(NAMESPACE)
	kubectl rollout restart deployment $(SALES_APP) --namespace=$(NAMESPACE)

dev-run: build dev-up dev-load dev-apply

dev-update: build dev-load dev-restart

dev-update-apply: build dev-load dev-apply

dev-logs:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(SALES_APP) --all-containers=true -f --tail=100 --max-log-requests=6 | go run api/tooling/logfmt/main.go -service=$(SALES_APP)

dev-logs-auth:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(AUTH_APP) --all-containers=true -f --tail=100 | go run api/tooling/logfmt/main.go

# ------------------------------------------------------------------------------
dev-logs-init:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(SALES_APP) -f --tail=100 -c init-migrate-seed

dev-describe-deployment:
	kubectl describe deployment --namespace=$(NAMESPACE) $(SALES_APP)

dev-describe-sales:
	kubectl describe pod --namespace=$(NAMESPACE) -l app=$(SALES_APP)

dev-describe-auth:
	kubectl describe pod --namespace=$(NAMESPACE) -l app=$(AUTH_APP)

# ==============================================================================
# Administration

migrate:
	export SALES_DB_HOST=localhost; go run api/tooling/admin/main.go migrate

seed: migrate
	export SALES_DB_HOST=localhost; go run api/tooling/admin/main.go seed

pgcli:
	pgcli postgresql://postgres:postgres@localhost

# ==============================================================================
# Metrics and Tracing

metrics-view:
	expvarmon -ports="localhost:3010" -vars="build,requests,goroutines,errors,panics,mem:memstats.HeapAlloc,mem:memstats.HeapSys,mem:memstats.Sys"

statsviz:
	open http://localhost:3010/debug/statsviz

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor

# ==============================================================================
# Running tests within the local computer

test-down:
	docker stop servicetest
	docker rm servicetest -v

test-r:
	CGO_ENABLED=1 go test -race -count=1 ./...

test-only:
	CGO_ENABLED=0 go test -count=1 ./...

lint:
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...

vuln-check:
	govulncheck ./...

test: test-only lint vuln-check

test-race: test-r lint vuln-check