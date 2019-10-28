pb:
	protoc -I . --go_out=plugins=micro:. proto/user/user.proto

build: pb
	GOOS=linux GOARCH=amd64 go build

run:
	docker build -t micro-user-service .
	docker run -p 50053:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		-e DB_HOST=localhost \
		-e DB_NAME=postgres \
		-e DB_USER=postgres \
		-e DB_PASSWORD=micro \
		micro-user-service