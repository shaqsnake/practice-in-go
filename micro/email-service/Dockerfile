FROM alpine:latest
# FROM debian:latest

RUN mkdir -p /app
WORKDIR /app

ADD email-service /app/email-service

CMD ["./email-service"]
