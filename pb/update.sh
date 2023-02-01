#!/bin/bash

protoc --go_out=plugins=grpc:. --go_opt=Mhello.proto=../pb  hello.proto
