#!/usr/bin/env bash

protoc --proto_path=./proto --go_out=./proto proto/*.proto