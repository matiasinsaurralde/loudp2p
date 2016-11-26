#!/bin/sh
mkdir -p pb
cd proto && protoc -I. --go_out=plugins=grpc:../pb *.proto
