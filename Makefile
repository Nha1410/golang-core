MAKEFILE_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
DEVBOX := research/devbox
INDEVBOX := docker run --network host --rm -i -v $(MAKEFILE_DIR):/go/src $(DEVBOX)

.PHONY: oapi-codegen
oapi-codegen: oapi-model-gen oapi-server-gen oapi-spec-gen

.PHONY: oapi-model-gen
oapi-model-gen:
	mkdir -p ./app/server/generated
	$(INDEVBOX) oapi-codegen --config=./docs/config/model.yaml ./docs/openapi.yaml > ./app/server/generated/model.go

.PHONY: oapi-server-gen
oapi-server-gen:
	mkdir -p ./app/server/generated
	$(INDEVBOX) oapi-codegen --config=./docs/config/server.yaml ./docs/openapi.yaml > ./app/server/generated/server.go

.PHONY: oapi-spec-gen
oapi-spec-gen:
	mkdir -p ./app/server/generated
	$(INDEVBOX) oapi-codegen --config=./docs/config/spec.yaml ./docs/openapi.yaml > ./app/server/generated/spec.go

.PHONY: gorm-gen
gorm-gen:
	docker compose exec fiber_app go run ./cmd/gormGen/main.go

setup-devbox:
	docker build -t research/devbox -f docker/devbox/Dockerfile .

build:
	docker build -t research/app -f docker/Dockerfile .

run:
	docker-compose up
