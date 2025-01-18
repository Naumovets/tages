FROM golang:1.23.5-alpine3.21 AS builder

COPY ./storage/ /source/
WORKDIR /source/

RUN go mod download
RUN go build -o ./bin/server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /source/bin/server .
COPY --from=builder /source/tools ./tools

CMD [ "./server" ]