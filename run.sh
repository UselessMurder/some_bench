#!/bin/bash

echo "Generate 500mb"
go run ./proto_bench/make.go
echo "Done"

echo "proto go:"
go run ./proto_bench/bench.go
echo "proto python3:"
python3 ./proto_bench/bench.py

echo "hex go:"
go run ./hex_bench/bench.go
echo "hex python3:"
python3 ./hex_bench/bench.py

rm ./bin
