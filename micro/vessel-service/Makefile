build:
	protoc -I . --go_out=plugins=micro:. proto/vessel/vessel.proto
	GOOS=linux GOARCH=amd64 go build

run:
	docker build -t micro-vessel-service .
	docker run -p 50052:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		micro-vessel-service