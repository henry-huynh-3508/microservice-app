FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

ADD consignment.json /app/consignment.json
ADD cli-consignment /app/cli-consignment

CMD ["./cli-consignment"]