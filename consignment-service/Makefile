protoc:
	protoc -I. --go_out=plugins=micro:. ./proto-consignment/consignment.proto

build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .
run:
	docker run -p 2000:2000 \
		-e MICRO_SERVER_ADDRESS=:2000\
		-e MICRO_REGISTR=mdns consignment-service