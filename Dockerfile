FROM golang:1.17.2

COPY janken/proto /go/src/janken/proto
ADD janken/go.*  /go/src/janken

WORKDIR /go/src/janken/

RUN apt update -y \
  && apt install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 && \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

RUN go mod vendor

RUN  protoc --go_out=. \ 
  --go-grpc_out=require_unimplemented_servers=false:. \
  ./proto/janken.proto

EXPOSE 50051

CMD go run ./cmd/api
