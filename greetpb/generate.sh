#!/bin/bash
#command to compile the proto in golang source file
protoc greet.proto --go_out=plugins=grpc:.