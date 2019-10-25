build:
	protoc -I . --go_out=plugins=micro:. proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build

run:
	docker build -t micro-consignment-service .
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		micro-consignment-service
