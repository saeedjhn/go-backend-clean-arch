BASE_PATH = .
USERS_PATH = users

## gen/v1/proto: Generate v1/proto services
.PHONY: gen/v1/proto
gen/v1/proto:
	@echo "> generating all services"
	protoc -I ${BASE_PATH} \
	--go_out=. \
	--go-grpc_out=. \
	${BASE_PATH}/**/*.proto

## gen/users: Generate v1/proto/users service
.PHONY: gen/users
gen/users:
	@echo "> generating users service"
	protoc -I ${USERS_PATH} \
	--go_out=. \
	--go-grpc_out=. \
	${USERS_PATH}/*.proto \
