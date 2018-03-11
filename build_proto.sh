#!/usr/local/bin bash
protoc --proto_path=./proto --go_out=plugins=micro:./src/share/pb ./proto/*.proto
