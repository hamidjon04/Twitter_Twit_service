include .env
export $(shell sed 's/=.*//' .env)

CURRENT_DIR := $(shell pwd)
DB_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

MIGRATE_PATH := database/migrations
MIGRATE_CMD := migrate -path $(MIGRATE_PATH) -database '$(DB_URL)' -verbose

.PHONY: gen-proto mig-up mig-down mig-force mig-create

gen-proto:
	./scripts/gen-proto.sh $(CURRENT_DIR)

mig-up: 
	$(MIGRATE_CMD) up

mig-down:
	$(MIGRATE_CMD) down

mig-force:
	$(MIGRATE_CMD) force 1

mig-create:
	migrate create -ext sql -dir $(MIGRATE_PATH) -seq user_tables
