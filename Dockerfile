FROM golang:alpine

WORKDIR /build
COPY simple-web-service .

EXPOSE 8080
CMD ["/build/simple-web-service"]
