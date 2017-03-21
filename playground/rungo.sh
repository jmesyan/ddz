#!/bin/bash

gofmt -w $1
golint $1
go run $1
