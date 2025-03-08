BASE_PATH = ./proto
USER_PATH = user

## gen/v1/proto: Generate v1/proto services
.PHONY: gen/v1/proto
gen/v1/proto:
	@echo "> generating all services"
	protoc -I ${BASE_PATH} \
	--go_out=. \
	--go-grpc_out=. \
	${BASE_PATH}/**/*.proto

## gen/users: Generate v1/proto/user service
.PHONY: gen/user
gen/user:
	@echo "> generating user service"
	protoc -I ${USER_PATH} \
	--go_out=. \
	--go-grpc_out=. \
	${USER_PATH}/*.proto \
