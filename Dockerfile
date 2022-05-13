# build stage
FROM golang:1.15 AS builder
WORKDIR /go/src/APIGateway
COPY . .
RUN make install

# final stage
FROM ubuntu:20.04
WORKDIR /root/
COPY --from=builder /go/src/APIGateway/APIGateway .
EXPOSE 55688
ENTRYPOINT ["./APIGateway"]
