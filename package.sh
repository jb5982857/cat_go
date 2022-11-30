#!/bin/bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

scp ./main root@110.185.185.119:/root/go/cat_account

rm ./main