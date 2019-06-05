#!/usr/bin/env sh

protoc greeter.proto --go_out=plugins=grpc:.