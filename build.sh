#!/bin/bash
curl -L -o libduckdb.zip https://github.com/duckdb/duckdb/releases/download/v0.3.2/libduckdb-src.zip
unzip -f libduckdb.zip
# cgo doesn't like it if the cpp file is in same directory when compiling ¯\_(ツ)_/¯
mkdir -p src
mv duckdb.cpp src/
# Note: this takes quite a while to compile!
g++ -I. -fPIC -c -o duckdb.o src/duckdb.cpp
ar -rcs libduckdb.a duckdb.o
CGO_LDFLAGS="-L$(pwd) -static -pthread -lduckdb -lstdc++ -ldl -lm" CGO_CFLAGS="-I$(pwd)" go build
