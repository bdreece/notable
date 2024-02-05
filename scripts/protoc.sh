#!/bin/sh

find . -type f -name "*.proto" | xargs protoc \
    -I./api \
    --go_out=. \
    --go_opt="module=github.com/bdreece/notable" \
    --go-grpc_out=. \
    --go-grpc_opt="module=github.com/bdreece/notable" \
    --experimental_allow_proto3_optional
