CURRENT_DIR=$(shell pwd)

gen-proto:
	./scripts/gen-proto.sh ${CURRENT_DIR}

swag-init:
	swag init -g api/router.go -o api/docs

mod-tidy:
	go mod tidy