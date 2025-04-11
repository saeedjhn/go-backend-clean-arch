CURRENT_DIR := $(shell pwd)
USER_PROTO_DIR := $(CURRENT_DIR)/user

## gen/v1/protos: Generate v1/proto/** services
.PHONY: gen/v1/protos
gen/v1/protos:
	@echo "> generating all services"
	protoc -I ${CURRENT_DIR} \
	--go_out=. \
	--go-grpc_out=. \
	${CURRENT_DIR}/**/*.proto

## gen/user: Generate v1/proto/user service
.PHONY: gen/user
gen/user:
	@echo "> generating user service"
	protoc -I $(USER_PROTO_DIR) \
	--go_out=. \
	--go-grpc_out=. \
	$(USER_PROTO_DIR)/*.proto \

#.PHONY: gen/user
#gen/user:
#	@echo "> generating user service"
#	protoc -I ${USER_PROTO_PATH} \
#	--go_out=. \
#	--go-grpc_out=. \
#	${USER_PROTO_PATH}/*.proto \
