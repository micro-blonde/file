#!/bin/bash

mkdir ./out
rm -r file/* || mkdir -p ./file

protoc --go_out=./out --go_opt=paths=source_relative --go-grpc_out=./out\
    --go-grpc_opt=paths=source_relative -I . \
    template/properties.proto \
    template/file.proto \
    template/meta.proto \
    template/store.proto \
    template/service.proto

mv out/template/* ./file
rm -r ./out
